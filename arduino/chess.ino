#include "Stepper.h"
#include "Arduino.h"

const float STEPS_PER_REV = 32;
const float GEAR_RED = 64;
const float STEPS_PER_OUT_REV = STEPS_PER_REV * GEAR_RED;

//Pins entered in sequence 1-3-2-4
Stepper stepperMotor(STEPS_PER_REV, 8, 10, 9, 11);

void setup()
{
    Serial.begin(9600);
}

void loop()
{
    //int stepsRequired = STEPS_PER_OUT_REV;
    //stepperMotor.setSpeed(800);
    //stepperMotor.step(stepsRequired);
    //delay(500);
    //stepsRequired = -STEPS_PER_OUT_REV;
    //stepperMotor.setSpeed(700);
    //stepperMotor.step(stepsRequired);
    delay(500);
    Serial.println("hello");
}

