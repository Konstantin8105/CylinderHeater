package main

type coord struct {
	x, y float64
}

type triangle struct {
	p [3]coord
}

func (t triangle) area() float64 {
	return method2(t)
}

/*

double area_4node(Node na,Node nb, Node nc, Node nd)
{
    return 2*area_3node(na,nb,nc);//fabs(0.5*((nb.x - na.x)*(nc.y - na.y)- (nc.x - na.x)*(nb.y - na.y)))*2;
};


double Shape::Jx_node(Node n0, Node n1, Node n2)
{
    Node temp_n[3];
    temp_n[0] = n0;
    temp_n[1] = n1;
    temp_n[2] = n2;
//    double xc = (n0.x+n1.x+n2.x)/3.;
//    double yc = (n0.y+n1.y+n2.y)/3.;
    double a = area_3node(n0,n1,n2);
                //0.5*((n1.x - n0.x)*(n2.y - n0.y)-
                //     (n2.x - n0.x)*(n1.y - n0.y));
//    printf("a = %.4e\n",a);//DEBUG
    double X_MIN = temp_n[0].x;
    double Y_MIN = temp_n[0].y;
    for(int i=0;i<3;i++)
    {
        if(Y_MIN > temp_n[i].y)  Y_MIN = temp_n[i].y;
        if(X_MIN > temp_n[i].x)  X_MIN = temp_n[i].x;
    }

//    for(int i=0;i<3;i++)//DEBUG
//        temp_n[i].Printf();//DEBUG
    for(int i=0;i<3;i++)
    {
        temp_n[i].x -= X_MIN;
        temp_n[i].y -= Y_MIN;
    }
//    for(int i=0;i<3;i++)//DEBUG
//        temp_n[i].Printf();//DEBUG

//    printf("Y_MIN = %.4f\n",Y_MIN);//DEBUG

    type_LLU x_left = 0, x_mid = 0, x_right =0;
    if(temp_n[0].x >= temp_n[1].x && temp_n[0].x > temp_n[2].x)
    {
        x_right = 0;
        if(temp_n[1].x > temp_n[2].x) { x_mid = 1; x_left = 2;}
        else                          { x_mid = 2; x_left = 1;}
    };
    if(temp_n[1].x >= temp_n[0].x && temp_n[1].x > temp_n[2].x)
    {
        x_right = 1;
        if(temp_n[0].x > temp_n[2].x) { x_mid = 0; x_left = 2;}
        else                          { x_mid = 2; x_left = 0;}
    };
    if(temp_n[2].x >= temp_n[1].x && temp_n[2].x > temp_n[0].x)
    {
        x_right = 2;
        if(temp_n[0].x > temp_n[1].x) { x_mid = 0; x_left = 1;}
        else                          { x_mid = 1; x_left = 0;}
    };
    if(temp_n[x_left].x == temp_n[x_mid].x && temp_n[x_left].y < temp_n[x_mid].y)
    {
        type_LLU r = x_left;
        x_left  = x_mid;
        x_mid   = r;
    }
    if(temp_n[x_right].x == temp_n[x_mid].x && temp_n[x_right].y < temp_n[x_mid].y)
    {
        type_LLU r = x_right;
        x_right = x_mid;
        x_mid   = r;
    }

    type_LLU type  = 0;
    double y0 =  temp_n[x_left ].y + (temp_n[x_right].y-temp_n[x_left ].y)/
                (temp_n[x_right].x-temp_n[x_left ].x)*(temp_n[x_mid].x-temp_n[x_left].x);
    if(temp_n[x_mid].y < y0) type = 0;
    else type = 1;
//    printf("TYPE = %u\n",type);//DEBUG
//
//    printf("x_left = %u\n",x_left);//DEBUG
//    printf("x_mid  = %u\n",x_mid );//DEBUG
//    printf("x_right= %u\n",x_right);//DEBUG

    double jx = -1e30;
    double Jx_left_mid    = Jx_node(temp_n[x_left ], temp_n[x_mid  ]);
//    printf("Jx_left_mid   = %.4e\n",Jx_left_mid  );//DEBUG
    double Jx_mid_right   = Jx_node(temp_n[x_mid  ], temp_n[x_right]);
//    printf("Jx_mid_right  = %.4e\n",Jx_mid_right );//DEBUG
    double Jx_left_right  = Jx_node(temp_n[x_left ], temp_n[x_right]);
//    printf("Jx_left_right = %.4e\n",Jx_left_right);//DEBUG


    if(type == 0)
    {
//        printf("TYPE 0 inside\n");
        jx  = +Jx_left_right
              -Jx_left_mid
              -Jx_mid_right;
//        line();//DEBUG
//        printf("Jx_left_mid   = %.4e\n",Jx_left_mid  );//DEBUG
//        printf("Jx_mid_right  = %.4e\n",Jx_mid_right );//DEBUG
//        printf("Jx_left_right = %.4e\n",Jx_left_right);//DEBUG
//        printf("jx = %.4e\n ",jx);//DEBUG
    }

    if(type == 1)
    {
//        printf("TYPE 1 inside\n");//DEBUG
        jx  = -Jx_left_right
              +Jx_left_mid
              +Jx_mid_right;
//        line();//DEBUG
//        printf("Jx_left_mid   = %.4e\n",Jx_left_mid  );//DEBUG
//        printf("Jx_mid_right  = %.4e\n",Jx_mid_right );//DEBUG
//        printf("Jx_left_right = %.4e\n",Jx_left_right);//DEBUG
//        printf("jx = %.4e\n ",jx);//DEBUG
    }
//    printf("LEFT \t");temp_n[x_left ].Printf();//DEBUG
//    printf("MID  \t");temp_n[x_mid  ].Printf();//DEBUG
//    printf("RIGHT\t");temp_n[x_right].Printf();//DEBUG
//    printf("TYPE = %u\n",type);//DEBUG
//    printf("Jx_left_mid   = %.4e\n",Jx_left_mid  );//DEBUG
//    printf("Jx_mid_right  = %.4e\n",Jx_mid_right );//DEBUG
//    printf("Jx_left_right = %.4e\n",Jx_left_right);//DEBUG
//    printf("jx = %.4e\n ",jx);//DEBUG




    double YC = (temp_n[0].y+temp_n[1].y+temp_n[2].y)/3.;
//    printf("YC = %e\n",YC);//DEBUG
//    printf("Y_MIN = %e\n",Y_MIN);//DEBUG
    if(jx < 1e-10) jx = 0;
    if(jx <0) {print_name("jx is less NULL");printf("jx[%e]\n",jx);}
    if(a  <0) print_name("area is less NULL");
    jx += -a*pow(YC,2.)+a*pow(Y_MIN+YC,2.);//*fabs(Y_MIN)/Y_MIN;

//    printf("[a*Y_MIN*Y_MIN] = %.4e\n jx = %.4e\n ",a*Y_MIN*Y_MIN,jx);//DEBUG
//    printf("jx = %.4e\n ",jx);//DEBUG
    return jx;
}

double Shape::Jx_node(Node n1, Node n2)
{
//    line();//DEBUG
//    printf("n1\t");//DEBUG
//    n1.Printf();//DEBUG
//    printf("n2\t");//DEBUG
//    n2.Printf();//DEBUG


    Node temp_n1 = n1;
    Node temp_n2 = n2;
    if(n1.y == 0 && n2.y == 0)
        return 0;
    if(n1.x == n2.x)
        return 0;
    if(n1.x == n2.x && n1.y == n2.y)
    {
        print_name("STRANGE");
        WARNING();
        return 0;
    }

    if(temp_n1.x > temp_n2.x)
    {
        swap(temp_n1.x,temp_n2.x);
        return Jx_node(temp_n1,temp_n2);
    }
    if(temp_n1.y > temp_n2.y)
    {
        swap(temp_n1.y,temp_n2.y);
        return Jx_node(temp_n1,temp_n2);
    }

    double jx = 0;
    double a = temp_n1.y;
    double b = fabs(temp_n2.x - temp_n1.x);
    double h = fabs(temp_n2.y - temp_n1.y);
    if(temp_n2.y < a) print_name("WARNING: position a");
//    if(temp_n1.y == temp_n2.y)
//    {
//        jx = (b*pow(a,3.)/12.+a*b*pow(a/2.,2.));
//        printf("**A**\n");
//    }
//    else
//    {
//    printf("a = %.4e\n",a);//DEBUG
//    printf("b = %.4e\n",b);//DEBUG
//    printf("h = %.4e\n",h);//DEBUG
//    printf("b*pow(a,3.)/12.    = %.4e\n",b*pow(a,3.)/12.);//DEBUG
//    printf("a*b*pow(a/2.,2.)   = %.4e\n",a*b*pow(a/2.,2.));//DEBUG
//    printf("b*pow(h,3.)/4.     = %.4e\n",b*pow(h,3.)/4.);//DEBUG
//    printf("(b*h/2.)*pow(a,2.) = %.4e\n",(b*h/2.)*pow(a,2.));//DEBUG
    jx = (b*pow(a,3.)/12.+a*b*pow(a/2.,2.))+(b*pow(h,3.)/12.+(b*h/2.)*pow(a,2.));
//    printf("jx = %.4e\n",jx);//DEBUG
    return fabs(jx);
};

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
