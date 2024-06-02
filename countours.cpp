#include "opencv2/imgproc.hpp"
#include "opencv2/imgcodecs.hpp"
#include "opencv2/highgui.hpp"
#include "opencv2/core.hpp"
#include <iostream>
#include <stdlib.h>

using namespace std;
using namespace cv;






Mat orig,im,vals,element,_red,_green;
Mat hsv,imgThreshold,final,gray,res,imgThreshold0,imgThreshold1,imgThreshold2;

static void help()
{
    cout << "\nThis program recognize clouds on satellite images"
        "\nCall:\n"
        "./cloud <image_name>\n"
        "\n"
        << endl;
}

int main(int argc, char **argv)
{
    cv::CommandLineParser parser(argc, argv, "{@input | ./data/lena.jpg | input image}");
    help();
    string input_image = parser.get<string>("@input");
    if (input_image.empty())
    {
        parser.printMessage();
        parser.printErrors();
        return 0;
    }

    Mat orig = imread(input_image);
/*
 // im = cv::CreateImage(orig.size(), 8, 1);
//CvtColor(orig, im, cv::CV_BGR2GRAY);
cv::cvtColor(orig,im, cv::COLOR_BGR2GRAY);

cv::threshold(im, im, 250, 255, cv::THRESH_BINARY);
imshow("Threshold 1", im);


int n=11;
element = getStructuringElement(MORPH_RECT, Size(n*2+1,n*2+1), Point(n, n));
cv::morphologyEx(im, im, cv::MORPH_OPEN, element);
cv::morphologyEx(im, im, cv::MORPH_CLOSE, element);

threshold(im, im, 200, 255, cv::THRESH_BINARY_INV);
imshow("After MorphologyEx", im);


//vals = im; #Make a clone because FindContours can modify the image
//contours=cv::FindContours(vals, cv::CreateMemStorage(0), cv::CV_RETR_LIST, cv::CV_CHAIN_APPROX_SIMPLE, (0,0));

vector<vector<Point> > contours;
findContours(im.clone(), contours, CV_RETR_LIST, CV_CHAIN_APPROX_SIMPLE);
_red = (0, 0, 255); 
_green = (0, 255, 0);
int levels=2 ;
//drawContours (orig, contours, _red, _green, levels, 2, cv::CV_FILLED); 

drawContours(orig, contours, -1, (0, 255, 0), 3);

imshow("Image", orig);
waitKey();
*/

 cvtColor( orig, gray, cv::COLOR_BGR2GRAY );
 imshow("gray",gray);
 

 

 //cv::inRange(gray,cv::Scalar(80,80,80), cv::Scalar(255,255,255), imgThreshold0);
 cv::inRange(gray,cv::Scalar(127,255,0), cv::Scalar(255,255,255), imgThreshold1);

 cv::inRange(gray,cv::Scalar(240,240,240), cv::Scalar(255,255,255), imgThreshold2);
 
 // imshow("thres1",imgThreshold0);
   imshow("thres1",imgThreshold1);
   imshow("thres2",imgThreshold2);
//bitwise_or(imgThreshold0,imgThreshold1,res);

vector<vector<Point> > contours; 
Mat hierarchy;




findContours(imgThreshold1, contours, hierarchy, RETR_CCOMP, CHAIN_APPROX_SIMPLE);
// draw contours 
Scalar color2 = Scalar(0,255, 0 );//green

for (int i = 0; i < contours.size(); ++i) {
    
    //drawContours(orig, contours, i, color, 1, LINE_8, hierarchy, 100);
   // drawContours(orig,i,0,color,-1);
    drawContours(orig,contours, i, color2, -1);
}


findContours(imgThreshold2, contours, hierarchy, RETR_CCOMP, CHAIN_APPROX_SIMPLE);
// draw contours 
Scalar color1 = Scalar(0,155, 0 );//darkgreen

for (int i = 0; i < contours.size(); ++i) {
    
    //drawContours(orig, contours, i, color, 1, LINE_8, hierarchy, 100);
   // drawContours(orig,i,0,color,-1);
    drawContours(orig,contours, i, color1, -1);
}



    
  // imshow("mask",res);
   imshow("out",orig);

waitKey();

}
