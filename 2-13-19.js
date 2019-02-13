function setup() {
  createCanvas(800, 800);
  background(255)
  strokeWeight(0.3)
  angleMode(DEGREES);
  drawSquare(0, 0, width, height, 0);
  
}

const maxDepth = 6;
const branchFactor = 3;

function drawSquare(x, y, dx, dy, depth) {
  if (depth >= maxDepth) {
   	rect(x, y, dx, dy);
    return;
  }
  
  if(Math.random() < 1-depth/maxDepth) {
    let localBranchFactor = branchFactor + (Math.random() * 2 - 1)
    let tileWidth = dx / localBranchFactor;
    let tileHeight = dy / localBranchFactor;
    for(let i = 0; i < localBranchFactor * localBranchFactor; i++) {
      let tx = Math.floor(i/localBranchFactor);
      let ty = i % localBranchFactor;
      drawSquare(x + tx * tileWidth, y + ty*tileWidth, tileWidth, tileHeight, depth+1);
    }
  } else {
    push()
    translate(x, y)
    rotate((5/2) - (Math.random() * 5))
    fill((depth/maxDepth * 255)+50 + Math.random() * 75)
   	rect(0, 0, dx, dy); 
    pop()
  }
}
