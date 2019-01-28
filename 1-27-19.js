var points;
var randFactor;
var yDelta;

function setup() {
	createCanvas(750, 750);
	background(Math.random() * 20 + 50);

	points = [];
	for (i = 0; i < 100; i++) {
		points.push(0);
	}

	randFactor = 1 + Math.random() * 3;
	frameRate(Math.random() * 5 + 7.5);
	yDelta = 10 + Math.random() * 10;
}

var y = 0;
function draw() {
	noFill();
	beginShape();
	colorMode(RGB, height, height, height, 1);
	strokeWeight(2.5);
	stroke(height - y, y, y / 2, 0.9);


	points = points.map((y_val, idx) => {
		curveVertex(idx * width / points.length, y_val + y);
		return y_val + Math.random() * randFactor
	});
	y += yDelta
	endShape();

	if (y > height) {
		clear();
		setup();
		y = 0;
	}
}
