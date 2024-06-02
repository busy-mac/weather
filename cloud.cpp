

#include "opencv2/imgproc.hpp"
#include "opencv2/imgcodecs.hpp"
#include "opencv2/highgui.hpp"
#include "opencv2/core.hpp"
#include <iostream>
#include <stdlib.h>

using namespace std;
using namespace cv;



const int BINARY_TH = 35;
const double SAT_FIX_TH = 0.3;
const double NBRR_FIX_TH = 0.4;




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

    Mat src = imread(input_image,IMREAD_COLOR);

    


cv::Mat ch[3], img_gray, img_element, img_brdiff,img_final;


img_element = getStructuringElement(MORPH_RECT, Size(6, 6), Point(5, 5));



img_brdiff = cv::Mat::zeros(src.size(), CV_8UC3);
cv::split(src, ch);
cv::absdiff(ch[2], ch[0], img_brdiff);
cv::threshold(img_brdiff, img_gray, BINARY_TH, 255, CV_THRESH_BINARY_INV);      
cv::morphologyEx(img_gray, img_gray, cv::MORPH_CLOSE, img_element);

double r, b, g, nbrr, sat;


for(int y = 0; y < src.rows; y++)
{
    for(int x = 0; x < src.cols; x++)
    {
        b = src.at<Vec3b>(y,x)[0];
        g = src.at<Vec3b>(y,x)[1];
        r = src.at<Vec3b>(y,x)[2];

        nbrr = (b - r) / (b + r);
        sat = 1.0 - (std::min(b, std::min(g, r)) / std::max(b, std::max(g, r)));

        if( nbrr < NBRR_FIX_TH &&
            sat < SAT_FIX_TH
            )
            img_gray.at<uchar>(y, x) = (uchar)255;
    }
}



cv::medianBlur(img_gray, img_gray, 3);

src.copyTo(img_final, img_gray);


//namedWindow("Source", 1);
    imshow("Source", src);

  // namedWindow("Final", 1);
   imshow("Final", img_final);



 




waitKey();
    return 0;

}
