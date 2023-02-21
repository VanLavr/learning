// if ( (x * x + y * y) < 0.5 ) pixel = '@'; circle
// if ( x <= 0.25 && x >= -0.25 && y <= 0.25 && y >= -0.25 ) pixel = '@'; square
// if ( y <= 1.6 * x + 0.7 && y >= 0.27 && y <= -1.6 * x + 0.7 && x <= 0.25 && x >= -0.25 ) pixel = '@'; triangle


// screen[i * width + j] = pixel // filling full screen
#include <stdio.h>
#include <math.h>
#include <Windows.h>

int main(int argc, char const *argv[]) {

    char pixel = ' ';

    const int width = 120;
    const int height = 30;
    const int screenLenght = width * height + 1;
    char screen[3601];
    screen[width * height] = '\0';

    float borderRelation = (float)width / (float)height; // relation width to height ( 4 : 1 )
    float symbolFrameRelation = 11.0f / 24.0f; // float symbolFrameRelation = 8.0f / 16.0f


	for (int k = 0; k < 10000; k++) {

		for (int i = 0; i < height; i++) {
			for (int j = 0; j < width; j++) {
				pixel = ' ';
				
				float x = 2 * ( ( (float)j / (float)width ) - 0.5f ); // normalizing width (from -1 to 1)
				float y = 2 * ( ( (float)i / (float)height ) - 0.5f ); // normalizing height (from -1 to 1)
				x *= borderRelation * symbolFrameRelation; // x and y coordinates in consider with screen and symbol border relations
                x += sin( (k * 0.01) - 0.5 );
                y += sin( (k * 0.015) - 0.85 );

                //y /= sin( (k * 0.01) - 0.5 );
                x /= sin( (k * 0.015) - 0.85 );

                if ( x <= 0.25 && x >= -0.25 && y <= 0.25 && y >= -0.25 ) pixel = '@';

				screen[i * width + j] = pixel;
			}
		}

        printf(screen);


        Sleep(3);
	}

    return 0;
}