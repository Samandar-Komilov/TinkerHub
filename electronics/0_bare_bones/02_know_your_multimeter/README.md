# 02 - Know Your Multimeter

## Goal
Make every dial position on your DT9205A mean something to you — not just the two you used in project 01.

## What you need
- DT9205A digital multimeter
- 9V battery + snap connector
- 1K resistor
- LED (any color)
- BC547 NPN transistor
- Jumper wires

## What to do

### Step 1: Measure DC voltage (DCV)

Turn the dial to **DCV 20V** (the range that covers up to 20V).
- Red probe in VΩmA, black probe in COM
- Touch red to battery +, black to battery -
- Read the voltage — should be somewhere around 7-9V depending on how fresh your battery is

**Now try the wrong range on purpose:**
- Switch to DCV 2V range and measure the 9V battery — you'll see "1" or "OL" (overload). The range is too small.
- Switch to DCV 200V — it reads, but with less precision (fewer decimal places)

**Takeaway:** Always pick the smallest range that's bigger than what you expect to measure.

### Step 2: Measure resistance (Ohm)

Turn the dial to **Ohm 2K** (the range that covers up to 2000 ohms).
- Touch the two probes to each end of your 1K resistor
- Read the value — it should show approximately 1000 (or close to it, maybe 980-1020)

**Try these:**
- Measure the same 1K resistor on the 200 ohm range — overload, range too small
- Measure it on 200K range — reads "1.0" meaning 1.0K, less precision
- Touch the probes to your skin (hold one in each hand) — you'll read your body resistance, probably 100K-2M depending on how dry your hands are

### Step 3: Continuity mode

You used this in project 01. Turn to the **continuity/diode** symbol.
- Touch probes together — beeps (short circuit, ~0 ohms)
- Touch across your 1K resistor — may or may not beep depending on your meter's threshold
- Touch across a jumper wire — beeps (wire has near-zero resistance)

**The rule:** continuity beeps when resistance is very low (typically under 30-50 ohms).

### Step 4: Diode mode

Same dial position as continuity on the DT9205A (shared symbol).
- Place your LED: red probe to the longer leg (anode), black probe to shorter leg (cathode)
- The display shows the **forward voltage** — red LEDs typically show ~1.7-1.8V, green ~2.0-2.1V
- Reverse the probes — "OL" (the LED blocks current in reverse, like a one-way valve)

**This tells you two things:**
1. Which leg is anode and which is cathode
2. How much voltage the LED "eats" when it's on

### Step 5: Measure DC current (DCA)

**Important: rewire your probes!**
- Move the red probe to the **20A** or **mA** jack (check your meter — this is a different hole)
- Turn dial to **DCA 20m** (milliamp range)

Build a simple circuit: 9V → 1K resistor → LED → back to battery. But **break the circuit** and insert the multimeter in series:

```
9V+ → red probe → [meter] → black probe → 1K resistor → LED → 9V-
```

- Read the current — should be roughly 5-10mA depending on your LED color
- The meter must be IN the current path (series), not across it (parallel)

**Warning:** If you accidentally measure current across the battery (probes directly on + and -), you short-circuit through the meter's low-resistance current shunt. The fuse may blow. Don't do this.

### Step 6: Transistor tester (hFE)

Your DT9205A has a small socket labeled **hFE** with holes marked E, B, C for NPN and PNP.
- Take your BC547 transistor
- Identify the pins (flat side facing you: E, B, C from left to right)
- Insert it into the **NPN** side of the hFE socket, matching E/B/C
- Read the number — this is the DC current gain (beta/hFE)

Typical BC547 values: 100-300. This number means: if 1mA flows into the base, up to hFE milliamps can flow through the collector. You'll use this number in project 15 and beyond.

### Step 7: AC voltage (ACV) — just know it exists

Turn to **ACV**. This measures alternating current voltage (like wall outlets).
- You won't use this with batteries (batteries are DC)
- If you ever need to measure a transformer or wall adapter's AC output, this is the mode
- For now, just know it's there

## Summary of dial positions

| Dial Position | What it measures | When you'll use it |
|--------------|-----------------|-------------------|
| DCV (various ranges) | DC voltage across two points | Every single project |
| Ohm (various ranges) | Resistance of a component | Checking resistors, LDR, pots |
| Continuity/Diode | Connection or LED forward voltage | Debugging, identifying LED polarity |
| DCA (mA, A ranges) | Current flowing through a wire | Understanding how much current LEDs and transistors use |
| hFE socket | Transistor current gain | Transistor projects (15+) |
| ACV | AC voltage | Not needed for battery circuits |

## Key takeaway
The multimeter is your eyes into the invisible. Voltage, current, resistance, continuity, component testing — it does all of it. From now on, when something doesn't work, your first move is: **measure it**.

## Evidence
After completing this, fill in your readings:

```
9V battery voltage:      _____ V
1K resistor measured:    _____ ohms
LED forward voltage:     _____ V  (color: _____)
LED circuit current:     _____ mA
BC547 hFE gain:          _____
Body resistance:         _____ K ohms
```
