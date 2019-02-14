const MAX_POW_2 = 7;
const MAX_RADIUS = 22;
const PADDING_X = 20;
const PADDING_Y = 20;

function setup() {
  createCanvas(1100, 850);
  background(255);
  strokeWeight(1.5)
  
  let sampleX = 50;
  let sampleY = 50;
  let tileWidth = (width - PADDING_X) / sampleX;
  let tileHeight = (height - PADDING_Y) / sampleY;
  
  noiseDetail(1, 0.5);
  translate(PADDING_X/2 + tileWidth/2, PADDING_Y/2 + tileHeight/2)
  for(let x = 0; x < sampleX; x++) {
   	for(let y = 0; y < sampleY; y++) {
       let sampledNoise = noise(x * tileWidth, y * tileHeight)
       strokeWeight(sampledNoise * 2 + 0.7)
     	 drawDigit(
      	x * tileWidth,
      	y * tileHeight,
      	1+(100 *sampledNoise) % 32
       );
    }
  }
}

function drawHexagon(x, y, r) {
	beginShape();
  for(var i = 1; i < 7; i++){
    const theta = i * 2 * Math.PI / 6;
  	vertex(x + r * Math.sin(theta), y + r * Math.cos(theta));
  }
  vertex(x, y+r);
	endShape(CLOSE);
}


function drawDigit(x, y, n) {
  stroke(51);
  fill(245);
  for(let i = MAX_POW_2; i >= 0; i--) {
    let pow2 = Math.pow(2, i);
  	if(n >= pow2) {
    	n -= pow2;
      drawHexagon(x, y, 0.6*lerp(2, MAX_RADIUS, i / MAX_POW_2));
    }
  }
}

