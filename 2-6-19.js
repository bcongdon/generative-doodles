function getRandom(min, max) {
    return Math.random() * (max - min) + min;
}

function setup() {
  createCanvas(800, 1000);
  background(255)
  
  for(var i = 0; i < 1000; i++) {
    var x1 = getRandom(-width, 0);
    var y1 = getRandom(-2*height, 2*height);
    var x2 = getRandom(width, 2*width);
    var y2 = getRandom(-height, 2*height);
  	drawStrip(x1, y1, x2, y2); 
  }
}

const stripeWidth = 5;
function drawStrip(x1, y1, x2, y2) {
  var theta = atan(-(x2-x1)/(y2-y1))
  var dy = stripeWidth * sin(theta);
  var dx = stripeWidth * cos(theta);
      
  strokeWeight(0);
  for(var i = 0; i < 7; i++) {
    if(i % 2 == 0)
      fill(51);
    else
      fill(250);

    quad(
    	x1, y1,
      x1+dx, y1+dy,
      x2+dx, y2+dy,
      x2, y2);

  	x1 += dx;
  	y1 += dy;
  	x2 += dx;
  	y2 += dy;
  }
}
