# Challenge 11 - Kepler-452

It’s the year 2145; AI still hasn’t surpassed human intelligence, the P versus NP problem is still a major unsolved problem in computer science, and there aren't any practical quantum computers. On the other hand, there have been enormous advances in space exploration. And thanks to the discovery of wormhole equations we can now explore the universe at faster than light speed.

You are in the business of unobtanium mining on the moons of the Kepler-452 system, where the mineral is extremely abundant and lucrative. You have mines on every moon of the Kepler-452 system and a fleet of cargo ships on every planet in the system. The cargo ships collect the mined mineral from the moons and takes it to the base planets through a series of warp jumps from moon to moon and back to the planet. Because doing a warp jump inside a planet's atmosphere is too dangerous, traditional rockets are used to get the ships to escape the planetary gravity. Consequently, you need to wait for a window of good weather conditions, and when you see one, you have to announce your route in advance. At launch time, every moon has a certain amount of unobtanium ready to be picked up at the mining facilities. So, you need to optimize your route to carry as much mineral as possible. Some things to take into account are:

* Loads cannot be separated. Once you land on a moon you have to take all the load available. And every ship in your fleet has a different payload capacity.
* You cannot go back to a moon that you've already visited. And you don't need to visit all the moons.
* Moons in Kepler-452 don't have an atmosphere, so you can warp jump from near ground level. However, every subsequent warp jump should be exactly **6 hours** after you land on a moon, which is more than enough time to load the ship and for the antimatter engines to cool down from the previous jumps.
* In outer space, every ship has a limited distance range of warp jumps it can make on a single charge and you need to make sure the ship has enough energy available to get back to the base planet. You can't refuel on the moons.
* Warp jumps are instantaneous.
* You don't need to worry about the additional payload for each trip. The energy needed to go from point A to point B through a wormhole only depends on the distance between those points. Because physics.
* All moons in Kepler-452 have a prograde orbit, near zero orbital eccentricity, and near zero degrees inclination from the planet orbit. But they have different orbital periods.
* Planets rotate counter-clockwise from our point of view.

#### Input
The first line is an integer **N** specifying the number of cases. **N** cases follow. Each case describes the initial conditions at the moment just before your cargo ship makes its first warp jump on a specific planet in the system. It has the following format:

* A line with an integer **M** that indicates the number of moons the planet has.
* A line with **M** float numbers, **d0, d1, …, dM-1**, separated by spaces, that indicates the distance from the planet to each moon, expressed in 106 kilometres.
* A line with **M** float numbers, **r0, r1, …, rM-1**, separated by spaces, that indicates the initial positions of the moons, expressed in radians, in their orbits around the planet at the moment just before the first warp jump.
* A line with **M** float numbers, **t0, t1, …, tM-1**, separated by spaces, that indicates the orbital period measured in hours.
* A line with **M** integers **u0, u1, …, uM-1**, separated by spaces, that indicates the amount of unobtanium ready to pick up on each moon measured in kilograms.
* A line with an integer **C**, that indicates the payload capacity of the ship measured in kilograms.
* A line with a float number **R**, that indicates the initial ship range expressed in 106 kilometres.

#### Limits
> 2 ≤ M ≤ 15 \
> 0.1 ≤ di ≤ 10 \
> 0 ≤ ri ≤ 2π \
> 10 ≤ ti ≤ 500 \ 
> 100 ≤ ui ≤ 2000 \ 
> 103 ≤ C ≤ 104 \
> 1 ≤ R ≤ 10

#### Output
N lines with the weight in kilograms of unobtanium brought from each moon visited. Each line should have the string “Case #n: ” followed by the kilograms of the unobtanium separated by whitespace and sorted in ascending order or 'None' if you can not visit any moon

#### Sample input

> 2 \
> 2 \
> 2.0 2.5 \
> 0.0 3.14 \
> 12.0 100.0 \
> 4 5 \
> 20 \
> 6.0 \
> 4 \
> 0.3 0.4 0.5 0.6 \
> 0.15 0.18 1.15 1.6 \
> 28.8 216.0 27.0 432.0 \
> 1532 770 1250 1630 \
> 3330 \
> 2.0

#### Sample output
> Case #1: 4 5 \
> Case #2: 1532 1630

##### Explanation
In case #1, we can visit both moons thanks to their movement (after 6h they become close to each other).

For case #2, we include this approximate drawing that shows the initial positions of the moons:

![](https://contest.tuenti.net/resources/img/kepler.png)

With the initial moon positions, the ship range and payload capacity, the best choice is to visit moons **m0** and **m3**, and bring back 1532 and 1630 kilograms of unobtanium.