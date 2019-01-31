function setup() {
  createCanvas(800, 600);
  background(220);
  makeTriangle(0, 0, width, height, 0);
}

function makeTriangle(x, y, w, h, depth) {
  stroke(255, depth*30, 0);
	if(depth > 7) {
  	return; 
  }
  strokeWeight(3/depth);
  triangle(x, y+h, x+w/2, y, x+w, y+h);

  
  makeTriangle(x, y+h/2, w/2, h/2, depth+1);
  makeTriangle(x+w/4, y, w/2, h/2, depth+1);
  makeTriangle(x+w/2, y+h/2, w/2, h/2, depth+1);
}
