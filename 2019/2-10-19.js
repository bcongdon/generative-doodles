function setup() {

    createCanvas(800, 800);
    stroke(0, 25);
    noFill();

    var tileWidth = width / 2, tileHeight = height / 2

    var minX = width, minY = height, maxX = 0, maxY = 0
    for(var i = 0; i < 4; i++) {
        var tileX = Math.floor(i / 2), tileY = i % 2
        var bounds = drawShape(
            i * 10,
            tileX * tileWidth,
            tileY * tileHeight,
            width - tileX * tileWidth,
            height - tileY * tileHeight
        )
        minX = min(minX, bounds[0])
        minY = min(minY, bounds[1])
        maxX = max(maxX, bounds[2])
        minY = max(maxY, bounds[3])
    }

    for(var i = 0; i < 16; i++) {
        var tileX = Math.floor(i / 4), tileY = i % 4
        drawShape(
            i * 10,
            tileX * tileWidth,
            tileY * tileHeight,
            width - tileX * tileWidth,
            height - tileY * tileHeight,
            minX, minY, maxX, maxY
        )
    }
}

function drawShape(detail, xmin, ymin, w, h, noiseXMin, noiseYMin, noiseXMax, noiseYMax) {
    noiseDetail(detail)
    var minX = width, minY = height
    var maxX = 0, maxY = 0
    for(var t = 0; t < 0.75; t += 0.005) {
        var x1 = xmin + w * noise(t + 15);
        var x2 = xmin + w * noise(t + 25);
        var x3 = xmin + w * noise(t + 35);
        var x4 = xmin + w * noise(t + 45);
        var y1 = ymin + h * noise(t + 55);
        var y2 = ymin + h * noise(t + 65);
        var y3 = ymin + h * noise(t + 75);
        var y4 = ymin + h * noise(t + 85);

        bezier(
            map(x1, noiseXMin, noiseXMax, xmin, xmin+w, false),
            map(y1, noiseYMin, noiseYMax, ymin, ymin+h, false),
            map(x2, noiseXMin, noiseXMax, xmin, xmin+w, false),
            map(y2, noiseYMin, noiseYMax, ymin, ymin+h, false),
            map(x3, noiseXMin, noiseXMax, xmin, xmin+w, false),
            map(y3, noiseYMin, noiseYMax, ymin, ymin+h, false),
            map(x4, noiseXMin, noiseXMax, xmin, xmin+w, false),
            map(y4, noiseYMin, noiseYMax, ymin, ymin+h, false)
        );

        minX = min(minX, x1, x2, x3, x4)
        minY = min(minY, y1, y2, y3, y4)
        maxX = max(minY, y1, y2, y3, y4)
        maxY = max(minY, y1, y2, y3, y4)
    }
    return [minX, minY, maxX, maxY]
}
