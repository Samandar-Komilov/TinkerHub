# Bare Bones Electronics

No Arduino. Just you, a battery, components, and a multimeter (DT9205A).
Each project builds on the previous one. Do them in order.

## Components assumed

From your kit + extras:
- Breadboard, jumper wires
- 9V battery + snap connector
- Resistors: 220R, 330R, 1K, 4.7K, 10K
- LEDs: red, green, yellow (multiple of each), RGB LED
- Potentiometer (10K)
- Push buttons (tactile)
- BC547 NPN transistor (x2)
- Diodes (1N4007)
- Capacitors: ceramic (100nF) and electrolytic (10uF, 100uF)
- Photoresistor (LDR)
- Piezo buzzer
- Small DC motor

---

## Part 1 — Your Tools

### 01_know_your_breadboard
**Problem:** You have a breadboard full of holes and no idea which ones are connected.
Plug your multimeter into continuity mode. Probe pairs of holes. Map out which rows connect, which don't, and where the power rails run. Draw the internal connections on paper.

### 02_know_your_multimeter
**Problem:** You have a DT9205A with 20+ positions on the dial and no intuition for when to use which.
Measure your 9V battery voltage (DCV 20V range). Measure a 1K resistor (Ω 2K range). Use the diode mode on an LED — see the forward voltage. Use continuity to buzz a wire. Try the hFE socket — plug in your BC547, read the gain number. By the end, every dial position should mean something to you.

### 03_first_circuit
**Problem:** You have a battery, an LED, and a resistor. Make the LED light up.
Wire 9V → 220R → LED → back to battery. It lights. Reverse the LED — it doesn't. Measure the voltage across the resistor and across the LED separately. They should add up to ~9V. (You just discovered something — keep that in mind.)

---

## Part 2 — Resistor Circuits

### 04_resistors_in_series
**Problem:** You have a 9V battery but no 4.5V battery. Can two resistors give you 4.5V somewhere?
Put two equal resistors (1K + 1K) in series across 9V. Measure voltage across each — about 4.5V each. Now try 1K + 10K. The voltages aren't equal anymore. Why does the bigger resistor get more voltage?

### 05_resistors_in_parallel
**Problem:** You need a 500R resistor but you only have 1K resistors.
Put two 1K resistors in parallel. Measure total resistance — ~500R. Two paths for current = less total resistance. Try 1K parallel with 10K — the result is close to 1K, not 5.5K. The smaller resistor dominates.

### 06_three_leds_in_series
**Problem:** You want 3 LEDs lit from one 9V battery in series. The third LED is dim or dead.
Wire 3 LEDs in series with a 220R. Measure voltage across each LED (~2V each) and the resistor. They add up to 9V. Three LEDs eat ~6V, resistor eats the rest. Try 4 LEDs — the fourth barely lights because 4×2V = 8V, leaving almost nothing for the resistor. **This is Kirchhoff's Voltage Law:** voltage rises and drops around any loop sum to zero. You didn't need to memorize it — you just lived it.

### 07_current_splits_at_a_junction
**Problem:** You wire 2 LEDs in parallel (each with its own resistor). Does each LED get the full battery current?
Build it: 9V → junction → two branches (220R + LED each) → rejoin → battery. Measure current into the junction. Measure current through each branch. Branch currents add up to total. **This is Kirchhoff's Current Law:** current into a node equals current out. Now try different resistors in each branch (220R and 1K) — more current takes the easier path.

### 08_led_brightness_control
**Problem:** Your LED is too bright with 220R. How do you dim it without buying new parts?
Same LED, try 220R → 1K → 10K. Measure current each time. Brightness tracks current, not voltage. At 10K the LED is barely visible. Now you understand Ohm's law viscerally: I = V/R. More resistance = less current = dimmer LED.

### 09_voltage_divider
**Problem:** You need 3V to test something, but you only have a 9V battery.
Two resistors in series: 10K + 4.7K. Measure the voltage at the middle point. It should be roughly 9 × 4.7/(10+4.7) ≈ 2.87V. Swap the resistor positions — now it's ~6.1V. You can dial in any voltage you want by choosing the ratio. This is the most used pattern in all of electronics.

### 10_voltage_divider_lies
**Problem:** Your voltage divider says 4.5V but the moment you connect an LED to it, the voltage drops to ~2V. What happened?
Build a divider with two 10K resistors (expect 4.5V). Measure — correct. Now connect an LED + 220R as a load. Measure again — the voltage sags badly. The load acts as a resistor in parallel with the bottom half, ruining your ratio. **This is why voltage dividers are bad power supplies.** The fix? You need something that doesn't care about load. (Transistors and regulators solve this — you'll get there.)

---

## Part 3 — Interaction

### 11_potentiometer
**Problem:** You want to smoothly vary brightness, not swap resistors by hand.
Connect pot as a voltage divider: one end to 9V, other to GND, wiper to 220R + LED. Rotate — brightness changes smoothly. Measure the wiper voltage at different positions. The pot is a voltage divider with an infinite number of positions.

### 12_push_button
**Problem:** You want an LED to turn on only when you press a button. Simple, right?
Wire: 9V → button → 220R → LED → GND. Press — LED on. Release — LED off. Now try the button wired differently: one side to 9V, other side to... nowhere? The wire just floats. Measure the "output" — it reads random garbage. **Floating inputs are the #1 beginner trap.** Fix it with a pull-down resistor (10K to GND) so the wire has a defined state when the button is open.

### 13_ldr_light_meter
**Problem:** How bright is the room? Can you measure light?
Build a voltage divider: LDR on top, 10K on bottom. Measure the middle voltage. Cover the LDR — voltage changes. Point it at a lamp — changes again. You just built a light sensor. The LDR is a resistor whose value depends on light. Map out the voltage at different light levels.

### 14_buzzer_experiments
**Problem:** You want your circuit to make sound, not just light.
Connect piezo buzzer across 9V (through a 220R). It clicks when you connect/disconnect. If you have a passive buzzer, rapid manual tapping makes a crude tone. Try it through a push button — you built a doorbell. Sound = vibration = something turning on and off fast. (You'll make this automatic soon.)

---

## Part 4 — Transistors

### 15_transistor_as_switch
**Problem:** Your push button is tiny and far away. You want a weak signal to control a strong one.
BC547: Emitter to GND. Collector to 220R + LED to 9V. Base to 10K + push button to 9V. Press button — tiny current through base (~0.5mA) switches on a much larger collector current (~20mA). The transistor is a current amplifier. Measure both currents. The ratio is the hFE gain you read on the multimeter in project 02.

### 16_transistor_not_turning_on
**Problem:** You wired a transistor but the LED stays off. Or it's dim. Why?
Try connecting the base directly to 9V without a resistor — the transistor gets hot. It works but you're overdriving it. Now try a 1M base resistor — LED barely glows. Not enough base current. Calculate: you need ~20mA collector current, your BC547 has hFE ~200, so base needs 20/200 = 0.1mA minimum. With 9V supply: R_base = 9V/0.1mA = 90K max. Try 10K (comfortable margin) vs 100K (barely works) vs 1M (doesn't work). **Base resistor selection is not a guess — it's math.**

### 17_automatic_night_light
**Problem:** You want an LED that turns on when the room goes dark. No code, no Arduino, just components.
Combine your LDR voltage divider (project 13) with a transistor switch (project 15). LDR + 10K divider → feeds the transistor base. In bright light, LDR resistance is low → low voltage at base → transistor off → LED off. Cover the LDR → resistance rises → base voltage rises → transistor turns on → LED lights up. You just built a real product with 5 components.

### 18_transistor_dimmer
**Problem:** You want smooth brightness control but the pot alone can't drive a bright LED without getting warm.
Pot → middle wiper to transistor base (through 10K). Transistor → drives LED. The pot handles only the tiny base current, the transistor handles the big LED current. Rotate pot — LED dims smoothly. Measure: pot wiper current is microamps, LED current is milliamps. This is amplification in action.

### 19_motor_control
**Problem:** You want to spin a DC motor from 9V but it draws too much current for a direct connection to your signal wire.
Transistor as switch, but for the motor. 9V → motor → transistor collector. Base driven through 1K from push button. Press — motor spins. **Important:** when you release, measure a voltage spike at the collector — the motor's inductance kicks back. It can kill your transistor. (You'll fix this in project 22.)

### 20_touch_sensitive_led
**Problem:** Can you make an LED turn on by touching a wire?
If you have 2 BC547s: connect them as a Darlington pair (emitter of first → base of second). Collector of second drives LED. Leave the base of the first one as a bare wire. Touch it — your skin's resistance (~100K-1M) provides enough base current because the gain is β₁ × β₂ ≈ 40,000. LED lights from your finger. Let go — off. If you only have one transistor, touch the base wire directly — it might flicker dimly (single transistor gain isn't enough for reliable touch).

### 21_blinking_leds
**Problem:** You want LEDs to blink on their own with no code, no clock, just components. Is that even possible?
Build an astable multivibrator: 2 transistors, 2 capacitors (10-100uF), 2 LEDs with 220R, 2 crossover resistors (10K-47K). Power it — the LEDs alternate blinking. Change the cap values to change speed. **How it works:** each capacitor charges through a resistor until it turns on the opposite transistor, which turns off the current side. The cycle repeats forever. This is the simplest oscillator — a circuit that has NO stable state.

---

## Part 5 — Capacitors

### 22_capacitor_charge_and_discharge
**Problem:** You want an LED to fade out slowly after you disconnect the battery, not snap off instantly.
Charge a large electrolytic (100uF) from 9V through a 1K resistor. Measure voltage across the cap as it charges — it rises and slows down near 9V. Disconnect the battery, connect an LED + 220R across the cap. Watch the LED fade over a second or two as the cap drains. **A capacitor is a tiny rechargeable battery.**

### 23_rc_timing
**Problem:** From project 22 — how long does the fade last? Can you predict it?
Build an RC circuit: 10K + 100uF. τ = R × C = 10K × 100uF = 1 second. Charge fully, then discharge through the resistor. Time how long it takes to drop to ~63% of the original voltage (that's 1τ). Change to 47K — now it takes ~5 seconds. **You can design precise delays with just two components.** This is how every timer circuit starts.

### 24_button_debounce
**Problem:** Your push button from project 12 looks clean to your eye, but if you watch the voltage on the multimeter, it jitters during the press. One "press" could look like 5 presses to a fast circuit.
Put a small capacitor (100nF) across the button output + a 10K resistor. The cap smooths the jitter — it can't charge/discharge fast enough to follow the bouncing contact. Measure with and without the cap. This is called debouncing, and every real button circuit needs it.

### 25_motor_noise_filter
**Problem:** When your motor from project 19 runs, nearby circuits act strange. The motor is generating electrical noise.
Solder/connect a 100nF ceramic cap directly across the motor terminals. Run the motor again. The noise is reduced — the cap shorts out the high-frequency spikes from the brushes. This is decoupling, and it's why every circuit board is littered with small caps near anything that switches.

---

## Part 6 — Diodes

### 26_diode_one_way_valve
**Problem:** You want current to flow only one direction — protect your circuit from reversed battery.
Put a 1N4007 in series with your LED circuit. LED works. Reverse the diode — LED goes dark. Current flows anode → cathode only. Measure the voltage drop across the diode when conducting: ~0.7V for silicon. Now measure LED forward voltage in the same way — it's a diode too, just one that glows.

### 27_flyback_protection
**Problem:** Your motor circuit from project 19 has dangerous voltage spikes when turning off.
Add a 1N4007 across the motor (cathode to +, anode to −, reversed from normal). When the motor is running, the diode does nothing (reverse biased). When you switch off, the motor's collapsing magnetic field creates a reverse voltage spike — the diode clamps it safely. Measure at the collector with and without the diode. **Every inductive load needs this.**

### 28_led_colors_are_voltages
**Problem:** Why do different color LEDs have different brightness with the same resistor?
Measure forward voltage of each LED color using your multimeter's diode mode: red (~1.8V), yellow (~2.0V), green (~2.1V), blue (~3.0V), white (~3.2V). Higher forward voltage = higher energy photons = shorter wavelength. Blue/white LEDs drop so much voltage that series chains max out faster. This is why project 06's "3 LEDs in series" math depends on LED color.

### 29_reverse_polarity_protection
**Problem:** What if someone plugs the battery in backwards? Protect the entire circuit.
Put a 1N4007 at the power input of any circuit. Normal polarity — diode conducts, everything works (minus 0.7V). Reversed — diode blocks, nothing gets damaged. Build your night light from project 17 with this protection. Cost: 0.7V of your supply voltage. Worth it.

---

## Part 7 — Combining Everything

### 30_rgb_color_mixing
**Problem:** You want arbitrary LED colors, not just red/green/yellow.
RGB LED has 3 LEDs in one package. Each needs its own resistor. Connect all three with 220R each. All on = white. Red + green = yellow. Red + blue = magenta. Now replace fixed resistors with pot-controlled transistor dimmers (project 18 pattern, times three). You have full color control.

### 31_light_level_indicator
**Problem:** You want to show light intensity as a bar graph — dim light = 1 LED, bright light = 3 LEDs.
Use your LDR voltage divider to feed different transistor switches with different threshold resistors. In dim light, only the most sensitive transistor turns on (1 LED). Brighter → 2 transistors on → 2 LEDs. Full brightness → all 3. You've built an analog-to-discrete converter from bare components.

### 32_dark_knowledge
**Problem:** You've been careful this whole time. Now break things on purpose to understand failure.
- Skip the resistor on an LED — watch it flash bright and dim forever (or die). Measure current at the moment of connection.
- Leave a transistor base floating — watch the LED flicker randomly from noise.
- Reverse an electrolytic capacitor at low voltage — it gets warm (stop quickly). This is why polarity markings exist.
- Connect motor without flyback diode, switch off rapidly — see/measure the spike.
- Exceed the voltage divider's expected range — see predictions break.

Controlled failure = deep understanding.

---

After completing all 32 projects, you will be confident with:
- Breadboard wiring and multimeter proficiency (including hFE, diode mode, capacitance)
- Kirchhoff's voltage and current laws — from real problems, not textbooks
- Voltage dividers — including why they fail under load
- Resistors (series, parallel, pull-up/pull-down)
- Capacitors (timing, filtering, debouncing, decoupling)
- Diodes (rectification, protection, forward voltage)
- LEDs (current limiting, color physics, RGB mixing)
- Potentiometers and LDRs (variable resistance, sensing)
- Transistors (switching, amplification, Darlington, oscillators, motor driving)

You are ready for Arduino.

---

## Worth Trying — Components You'd Need to Buy

These projects solidify bare-bones knowledge but need components not in a typical starter kit.

| Project | What you'd build | Components needed |
|---------|-----------------|-------------------|
| **Thevenin in practice** | Reduce a complex resistor network to one equivalent source. Verify by measurement. | Assorted resistor pack (need 6+ different values to make it worthwhile) |
| **Wheatstone bridge** | Precision resistance measurement — more accurate than your multimeter's Ω mode. | 1x precision pot (multi-turn 10K) + 1% tolerance resistors |
| **Zener voltage regulator** | Get a stable 5.1V from 9V that doesn't sag under load like a voltage divider. | Zener diodes (5.1V, 3.3V) |
| **Full-wave bridge rectifier** | Convert AC to bumpy DC, then smooth it with a cap. See how power supplies work. | Small 9V AC transformer + 4x 1N4007 (you have diodes but need the transformer) |
| **Transistor amplifier** | Amplify a small audio signal from your phone's headphone jack into a speaker. | 8Ω small speaker + 3.5mm audio jack + 2x 10uF coupling caps |
| **H-bridge motor control** | Drive a motor forward AND reverse with transistors. | 2x PNP transistor (BC557) — you only have NPN |
| **Push-pull output stage** | NPN + PNP pair driving a load efficiently in both directions. | BC557 PNP transistor |
| **Constant current source** | LED brightness that doesn't change when battery drains from 9V to 7V. | LM317 voltage regulator (used as current source) or just a few extra transistors |
| **Op-amp comparator** | Clean switching with hysteresis — your night light but with a sharp on/off threshold instead of gradual. | LM358 or LM741 op-amp IC |
| **Op-amp amplifier** | Inverting and non-inverting amplifier with precise, predictable gain. | LM358 dual op-amp + a few extra resistors |
| **Op-amp oscillator** | Generate a steady square wave with precise frequency — cleaner than the transistor astable. | LM358 + a few caps and resistors |
| **Voltage follower** | Buffer that isolates your voltage divider from its load — fixes the problem from project 10 completely. | LM358 op-amp |
