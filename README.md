# gotracer
A simple concurrent CPU pathtracer written in vanilla Go. Based off the pathtracer outlined in the book "Ray Tracing in One Weekend" and following by Peter Shirley. 

---
![Examples](http://i.imgur.com/g9jVQzX.png)
---
A fairly simple application right now. Currently generates three large sphere surrounded by a field of smaller spheres of
various materials and colors. 

### Features
* Lambertian, metal, and dielectric materials
* Spheres
* Camera positioning
* Feild of view
* Depth of field
* Outputs to PNG file
* Anti-alliasing
* Command line arguments
* Multicore support
* Motion blur
* BVH tree for faster renderering
* Basic textures

### Future Addition
* OBJ file support
* Additional shapes
* Texture mapping
* Lighting
* Additonal materials
* Adaptive sampling
