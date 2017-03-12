package section

import (
	"math"
)

type sectionTriangles struct {
	triangles []triangle
}

func (s sectionTriangles) area() (area float64) {
	for _, tr := range s.triangles {
		if tr.check() == nil {
			area += tr.area()
		}
	}
	return
}

func (s sectionTriangles) centerMassX() float64 {
	var summs float64
	var areas float64
	for _, tr := range s.triangles {
		area := tr.area()
		summs += area * tr.centerMassX()
		areas += area
	}
	return summs / areas
}

func (s sectionTriangles) centerMassZ() float64 {
	var summs float64
	var areas float64
	for _, tr := range s.triangles {
		area := tr.area()
		summs += area * tr.centerMassZ()
		areas += area
	}
	return summs / areas
}

func (s sectionTriangles) momentInertiaX() (j float64) {
	zc := s.centerMassZ()
	for _, tr := range s.triangles {
		if tr.check() == nil {
			tm := triangle{[3]coord{
				coord{x: tr.p[0].x, z: tr.p[0].z - zc},
				coord{x: tr.p[1].x, z: tr.p[1].z - zc},
				coord{x: tr.p[2].x, z: tr.p[2].z - zc},
			}}
			j += tm.momentInertiaX()
		}
	}
	return
}

func (s sectionTriangles) momentInertiaZ() (j float64) {
	xc := s.centerMassX()
	for _, tr := range s.triangles {
		if tr.check() == nil {
			tm := triangle{[3]coord{
				coord{x: tr.p[0].x - xc, z: tr.p[0].z},
				coord{x: tr.p[1].x - xc, z: tr.p[1].z},
				coord{x: tr.p[2].x - xc, z: tr.p[2].z},
			}}
			j += tm.momentInertiaZ()
		}
	}
	return
}

func (s sectionTriangles) minimalMomentOfInertia() (j float64) {
	// degree 0
	Jxo := s.momentInertiaX()
	Jzo := s.momentInertiaZ()
	// degree 45
	alpha45 := 45. / 180. * math.Pi
	var rotateTriangle []triangle
	for _, tr := range s.triangles {
		var rTriangle triangle
		for i := range tr.p {
			lenght := math.Sqrt(tr.p[i].x*tr.p[i].x + tr.p[i].z*tr.p[i].z)
			alpha := math.Atan(tr.p[i].z / tr.p[i].x)
			alpha += alpha45
			rTriangle.p[i] = coord{
				x: lenght * math.Cos(alpha),
				z: lenght * math.Sin(alpha),
			}
		}
		rotateTriangle = append(rotateTriangle, rTriangle)
	}
	Jx45 := sectionTriangles{triangles: rotateTriangle}.momentInertiaX()

	// f = (cos45)^2 = (sin45)^2
	f := math.Pow(math.Cos(45./180.*math.Pi), 2.)
	Jxyo := Jxo*f - Jx45 + Jzo*f
	alpha := math.Atan(2 * Jxyo / (Jzo - Jxo))

	Ju := Jxo*math.Pow(math.Cos(alpha), 2.) - Jxyo*math.Sin(2*alpha) + Jzo*math.Pow(math.Sin(alpha), 2.)
	Jv := Jxo*math.Pow(math.Sin(alpha), 2.) + Jxyo*math.Sin(2*alpha) + Jzo*math.Pow(math.Cos(alpha), 2.)
	return math.Min(Ju, Jv)
}

func (s sectionTriangles) sectionModulusWx() (j float64) {
	var zmax float64
	zc := s.centerMassZ()
	for _, tr := range s.triangles {
		for _, c := range tr.p {
			zmax = math.Max(zmax, c.z-zc)
		}
	}
	return s.momentInertiaX() / zmax
}
func (s sectionTriangles) sectionModulusWz() (j float64) {
	var xmax float64
	xc := s.centerMassX()
	for _, tr := range s.triangles {
		for _, c := range tr.p {
			xmax = math.Max(xmax, c.x-xc)
		}
	}
	return s.momentInertiaZ() / xmax
}

func (s sectionTriangles) check() error {
	for _, tr := range s.triangles {
		if err := tr.check(); err != nil {
			return err
		}
	}
	return nil
}

/*
double Shape::CalcJ(double Angle)
{
    double J = 0;
    mesh->RotatePointXOY(0,0,Angle);
    for(type_LLU i=0;i<mesh->elements.GetSize();i++)
    {
        Element el = mesh->elements.Get(i);
        if(el.ElmType == ELEMENT_TYPE_TRIANGLE)
        {
            Node p[3];
            p[0] = mesh->nodes.Get(el.node[0]-1);
            p[1] = mesh->nodes.Get(el.node[1]-1);
            p[2] = mesh->nodes.Get(el.node[2]-1);
            J += Jx_node(p[0],p[1],p[2]);
        }
    }
    mesh->RotatePointXOY(0,0,-Angle);
    return J;
}

double Shape::AngleWithMinimumJ(double step0, double _angle)
{
    bool DEBUG = false;//true;
    double x0 = _angle-step0*1;
    double x1 = _angle+step0*0;
    double x2 = _angle+step0*1;
    double y0 = CalcJ(x0);
    double y1 = CalcJ(x1);
    double y2 = CalcJ(x2);
    double eps= 1e-6;
    if(DEBUG)printf("step = %.5e x0=%.5e x1=%.5e x2=%.5e\n",GRADIANS(step0),GRADIANS(x0),GRADIANS(x1),GRADIANS(x2));
    if(GRADIANS(max(x0,x1,x2)-min(x0,x1,x2))<=eps || max(y0,y1,y2)-min(y0,y1,y2)<=eps*min(y0,y1,y2))
    {
             if(y0 == min(y0,y1,y2)) return x0;
        else if(y1 == min(y0,y1,y2)) return x1;
        else return x2;
    }
         if(min(y0,y1,y2) == y0) return AngleWithMinimumJ(step0, x0);
    else if(min(y0,y1,y2) == y2) return AngleWithMinimumJ(step0, x2);
    else                         return AngleWithMinimumJ(step0/1.5, x1);
}

*/
