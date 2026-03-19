# Arduino Uno Projects

You've built bare-bones circuits. Now add a brain.
Each project builds on previous ones + your bare-bones knowledge.

## Projects

### 01_blink_onboard
**Learn:** Arduino IDE setup, uploading code, digital output.
Blink the built-in LED (pin 13). No wiring needed — pure software start.

### 02_blink_external
**Learn:** Connecting Arduino to breadboard, digitalOutput to external LED.
Blink an external LED on a breadboard with 220R resistor.

### 03_traffic_light
**Learn:** Multiple digital outputs, timing/sequencing logic.
3 LEDs (red/yellow/green) cycling like a traffic light.

### 04_button_input
**Learn:** digitalRead, pull-up/pull-down resistors, input vs output.
Press a button to turn LED on/off. Understand debouncing.

### 05_analog_read_pot
**Learn:** analogRead, ADC (0-1023), serial monitor.
Read potentiometer value, print to serial monitor. See analog world become digital.

### 06_pwm_led_fade
**Learn:** analogWrite (PWM), duty cycle, smooth LED dimming via code.
Fade LED in and out smoothly. Compare to bare-bones pot dimming.

### 07_pot_controls_led
**Learn:** Combine analog input + PWM output.
Potentiometer controls LED brightness through Arduino. Input -> processing -> output.

### 08_serial_communication
**Learn:** Serial.print, reading sensor data, debugging with serial.
Send messages between Arduino and computer. Foundation for all future debugging.

### 09_transistor_motor_driver
**Learn:** Use transistor (BC547) to drive loads Arduino can't power directly.
Arduino signal -> transistor -> LED strip or small motor from 9V battery.

### 10_multi_sensor_dashboard
**Learn:** Combine multiple inputs, serial output, real-time monitoring.
Pot + button(s) -> Arduino reads all, reports state via serial. Your first "system."

---

After completing all 10 projects, you will be confident with:
- Arduino IDE and uploading
- Digital I/O (read/write)
- Analog input (ADC)
- PWM output (fake analog)
- Serial communication and debugging
- Combining bare-bones knowledge (resistors, transistors, pots) with code
- Driving external loads through transistors
- Building multi-component systems

You are ready for sensors, displays, and real projects.
