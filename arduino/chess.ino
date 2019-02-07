#include "Stepper.h"
#include "Arduino.h"

const float STEPS_PER_REV = 32;
const float GEAR_RED = 64;
const float STEPS_PER_OUT_REV = STEPS_PER_REV * GEAR_RED;
const float REV_PER_CM = 0.5;
const float maxSpeed = 700;

//Pins entered in sequence 1-3-2-4
Stepper stepperMotor(STEPS_PER_REV, 8, 10, 9, 11);
char inData[20];
// Position in cm
float posX;
float posY;

void setup()
{
    Serial.begin(9600);
    posX = 0;
    posY = 0;
}

float extractNumber(char *data, int n, int *end)
{
    *end = 0;
    while (*end < n && data[*end] != '/' && data[*end] != '\0') (*end)++;
    data[*end] = '\0';
    return atof(data);
}

void receivePosition(float *x, float *y)
{
    char inChar = '0';
    int index = 0;
    do
    {
        if (Serial.available() > 0) {
            inChar = Serial.read();
            inData[index] = inChar;
            index++;
        }
    } while (inChar != '\n' && index < 20);
    inData[index-1] = '\0';
    int end;
    Serial.println(inData);
    *x = extractNumber(inData, index, &end);
    *y = extractNumber(inData+end+1, index, &end);
}

int incomingByte = 0;

void computeMove(float x, float y, float *moveX, float *moveY)
{
    *moveX = x - posX;
    *moveY = y - posY;
}

void move(float x, float y)
{
    float m = max(abs(x), abs(y));
    if (m == 0)
    {
        Serial.println("Not moving");
        return;
    }
    long speedX = long(maxSpeed * abs(x) / m);
    long speedY = long(maxSpeed * abs(y) / m);
    int stepsX = int(x*REV_PER_CM*STEPS_PER_OUT_REV);
    int stepsY = int(y*REV_PER_CM*STEPS_PER_OUT_REV);
}

void loop()
{
    float x,y;
    receivePosition(&x, &y);
    float moveX, moveY;
    computeMove(x, y, &moveX, &moveY);
    //Move
    posX = x;
    posY = y;
    //Serial.print("X is ");
    //Serial.println(x);
    //Serial.print("Y is ");
    //Serial.println(y);
    //int stepsRequired = STEPS_PER_OUT_REV;
    //stepperMotor.setSpeed(800);
    //stepperMotor.step(stepsRequired);
    //delay(500);
    //stepsRequired = -STEPS_PER_OUT_REV;
    //stepperMotor.setSpeed(700);
    //stepperMotor.step(stepsRequired);
    //delay(500);
}

