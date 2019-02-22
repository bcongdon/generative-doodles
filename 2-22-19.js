function setup() {
  createCanvas(1000, 800);
  background(225)
  fill(245)
  translate(width/200, height/200)
  scale(0.99, 0.99)
  let stroke = 1;
  for(let i = 0; i < 3000; i++) {
    rect(0, 0, width-1, height-1)
    translate(i/2, sin(i/100)+i/2)
    scale(0.965, .975)
    strokeWeight(stroke)
    stroke /= 0.978
  }
}
