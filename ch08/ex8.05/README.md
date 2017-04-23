# Exercise 8.5 (P239)

Take an existing CPU-bound sequential program, such as the Mandelbrot program of Section 3.3 or the 3-D surface computation of Section 3.2,
and execute its main loop in parallel using channels for communication.
How much faster does it run on a multiprocessor machine?
What is the optimal number of goroutines to use?

## Results

Render 4096*4096 Mandelbrot

``` text
    NumCPU: 4
    Done no concurrency. Used: 5.144318122s
    Done. Worker Number: 1 Used: 5.060132177s
    Done. Worker Number: 2 Used: 2.820846794s
    Done. Worker Number: 3 Used: 2.781336072s
    Done. Worker Number: 4 Used: 2.579062955s
```
