# 01 - Know Your Breadboard

## Goal
Understand how your breadboard is wired internally — before plugging anything in.

## What you need
- Breadboard (830-point)
- Digital multimeter
- 2 jumper wires

## What to do

### Step 1: Look at the board
Your 830-point breadboard has:
- **2 power rails** (top and bottom) — the long horizontal rows marked + and -
- **2 terminal blocks** (left and right halves) — the numbered vertical columns
- **A center gap** — splits the terminal blocks in half

### Step 2: Set your multimeter to continuity mode
- Turn the dial to the continuity/beep symbol (looks like a sound wave or diode symbol)
- Touch the two probes together — it should beep. That means "connected"

### Step 3: Test which holes connect to each other

**Test these and write down what beeps (connected) vs what doesn't:**

1. Two holes in the SAME row on one side of the center gap → beeps, random numbers shown
2. Two holes in the same row but ACROSS the center gap → nothing happened
3. Two holes in the SAME power rail (+ row, horizontal) → beeps (same + power rail)
4. The + rail on the TOP vs the + rail on the BOTTOM → nothing happened as center gap exists between
5. A power rail hole vs a terminal block hole in the same column → nothing happened

### Step 4: Draw it out
On paper, draw your breadboard and mark which holes connect to which.
This mental model is everything — every future circuit depends on it.

## What you should discover
- Each short row (5 holes) on one side of the gap is connected internally
- Rows do NOT connect across the center gap
- Power rails run horizontally along the full length (or sometimes split in the middle — test it!)
- Power rails do NOT connect to terminal block rows
- On some boards, top and bottom power rails are NOT connected to each other

## Key takeaway
The breadboard is just a bunch of hidden wires. Your multimeter proved it. Now you know exactly where current can and can't flow.

## Evidence
After completing this, note your findings here:

```
Same row, same side:     connected
Same row, across gap:    not connected
Same power rail:         connected
Top + rail vs bottom +:  not connected
Power rail vs terminal:  not connected
```
