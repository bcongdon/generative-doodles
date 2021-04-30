import vsketch
import math

class SpiralTunnelSketch(vsketch.SketchClass):
    # Sketch parameters:
    ticks_per_circle = vsketch.Param(25)
    compression = vsketch.Param(1.00)
    num_ticks = vsketch.Param(4200)
    mode = vsketch.Param("line", choices=["line", "triangle"])

    def draw(self, vsk: vsketch.Vsketch) -> None:
        vsk.size("8.5x11in", landscape=False)
        vsk.scale("cm")

        for idx in range(self.num_ticks):
            r1 = idx / 100.0 * self.compression + 0.25
            r2 = r1 + 0.025

            ticks_per_circle = self.ticks_per_circle
            angle = (360 / ticks_per_circle) * ((float(idx) % ticks_per_circle) + 1)* math.pi / 180.0

            vsk.line(
                math.cos(angle) * r1, math.sin(angle) * r1,
                math.cos(angle) * r2, math.sin(angle) * r2,
            )

        for i in range(3, 10):
            vsk.circle(0, 0, i**2/3.96)

        vsk.point(0, 0)
        
    def finalize(self, vsk: vsketch.Vsketch) -> None:
        vsk.vpype("linemerge linesimplify reloop linesort")


if __name__ == "__main__":
    SpiralTunnelSketch.display()
