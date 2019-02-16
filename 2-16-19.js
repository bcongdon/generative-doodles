function randInt(min, max) {
  return Math.random() * (max - min) + min
}

function setup() {
  createCanvas(1200, 800);
  background(250)
  angleMode(DEGREES);
  strokeWeight(0.15)
  
  let buckets = {}
  
  for(let i = 0; i < 1000; i++) {
    let x = Math.random() * width;
    let y = Math.random() * height;
    
    fill(255-noise(x, y)*50)
    stroke(noise(x, y)*255)
    
      push()
    translate(x, y)
    rotate(noise(x, y, 3) * 180)
       ellipse(0, 0, noise(x, y, 1) * 30, noise(x, y, 2) * 30);
    pop()
    
    let xBucket = Math.floor(x/(width/10))
    let yBucket = Math.floor(y/(height/8))
    if(!buckets[xBucket + " " + yBucket]) {
      buckets[xBucket + " " + yBucket] = []
    }
    
    buckets[xBucket + " " + yBucket].push([x, y])
  }
  
  stroke(100)
  for(let key in buckets) {
    let bucket = buckets[key]
    let i = 0
    bucket.forEach(coord1 => {
      bucket.forEach(coord2 => {
        if (i > 5)
          return
        if(Math.random() > 0.9) {
          strokeWeight(1/i + 0.7)
          beginShape();
          noFill();
          vertex(coord1[0], coord1[1]);
          quadraticVertex(
            coord2[0] + randInt(-5, 5),
            coord2[1] + randInt(-25, 25),
            coord2[0], coord2[1])
          endShape();
          i += 1
        }
      })
    })
  }
}

function draw() {
}
