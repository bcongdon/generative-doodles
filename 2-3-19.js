const MAX_POW_2 = 7;
const MAX_RADIUS = 22;

function setup() {
  createCanvas(475, 750);
  background(50);
  
  const padding_x = (width - 7 * (MAX_RADIUS * 2)) / 2;
  const padding_y = (height - (((Math.pow(2, MAX_POW_2) / 8) - 1) * (MAX_RADIUS * 2))) / 2;
  for(let i = 0; i < Math.pow(2, MAX_POW_2); i++) {
    drawDigit(padding_x + i % 8 * MAX_RADIUS * 2, padding_y + Math.floor(i / 8) * MAX_RADIUS * 2, i);
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
      drawHexagon(x, y, lerp(2, MAX_RADIUS, i / MAX_POW_2));
    }
  }
}

