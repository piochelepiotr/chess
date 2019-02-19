#include "Stepper.h"
#include "Arduino.h"
#include "AccelStepper.h"
#include "MultiStepper.h"

const float STEPS_PER_REV = 32;
const float GEAR_RED = 64;
const float STEPS_PER_OUT_REV = STEPS_PER_REV * GEAR_RED;
const float REV_PER_CM = 0.5;
const float STEPS_PER_CM = REV_PER_CM * STEPS_PER_OUT_REV;
const float maxSpeed = 700;

//Pins entered in sequence 1-3-2-4
AccelStepper stepperMotorY(AccelStepper::BYJ, 8, 9, 10, 11);
AccelStepper stepperMotorX(AccelStepper::BYJ, 4, 5, 6, 7);
MultiStepper steppers;

char inData[20];

void setup()
{
    Serial.begin(9600);
    stepperMotorX.setMaxSpeed(700);
    stepperMotorY.setMaxSpeed(700);
    steppers.addStepper(stepperMotorX);
    steppers.addStepper(stepperMotorY);
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
    //Serial.println(inData);
    *x = extractNumber(inData, index, &end);
    *y = extractNumber(inData+end+1, index, &end);
}

int incomingByte = 0;

void move(float x, float y)
{
    long positions[2];
    positions[0] = long(x*STEPS_PER_CM);
    positions[1] = long(y*STEPS_PER_CM);
    Serial.print("Move x to ");
    Serial.println(positions[0]);
    Serial.print("Move y to ");
    Serial.println(positions[1]);
    stepperMotorX.moveTo(positions[0]);
    steppers.moveTo(positions);
    steppers.runSpeedToPosition();
    Serial.println("finished");
    Serial.println("useless");
}

void loop()
{
    float x,y;
    receivePosition(&x, &y);
    move(x, y);
}

