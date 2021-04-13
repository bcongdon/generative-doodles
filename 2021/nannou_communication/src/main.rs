use nannou::prelude::*;

fn main() {
    nannou::app(model).update(update).run();
}

struct Model {
    things: Vec<Thing>,
    projectiles: Vec<Projectile>,
}

#[derive(Clone, Copy)]
struct Thing {
    position: Vector2,
    last_hit: f32,
    color: Hsl,
}

impl Thing {
    pub fn new(p: Vector2) -> Self {
        Thing {
            position: p,
            last_hit: 0.0,
            color: hsl(random_f32(), 1.0, 0.5),
        }
    }
}

struct Projectile {
    source: usize,
    color: Hsl,
    dest: usize,
    position: Vector2,
}

const N_THINGS: usize = 100;

fn model(app: &App) -> Model {
    let _window = app
        .new_window()
        .size(1024, 1024)
        .view(view)
        .build()
        .unwrap();
    let mut things = Vec::new();

    for _ in 0..N_THINGS {
        things.push(Thing::new(vec2(
            (random::<f32>() - 0.5) * 1024.0,
            (random::<f32>() - 0.5) * 1024.0,
        )));
    }

    Model {
        things: things.iter().cloned().collect(),
        projectiles: (0..10)
            .map(|_| {
                let dest = random_range(0, things.len());
                Projectile {
                    source: random_range(0, things.len()),
                    color: things[0].color,
                    dest: dest,
                    position: things[dest].position,
                }
            })
            .collect(),
    }
}

fn update(app: &App, model: &mut Model, _update: Update) {
    let time = app.elapsed_frames() as f32 / 60.0;

    let things = &mut model.things;
    let mut things_to_emit = Vec::new();
    model.projectiles.retain(|p| {
        let target = things[p.dest].position;
        let collided = (target - p.position).magnitude() < 2.0;
        if collided {
            things_to_emit.push(p.dest);
            things[p.dest].last_hit = time;
        }
        !collided
    });

    for idx in things_to_emit {
        if model.projectiles.len() > 100 {
            break;
        }
        for _ in 0..random_range(3, 50) {
            model.projectiles.push(Projectile {
                source: idx,
                dest: random_range(0, things.len()),
                position: things[idx].position,
                color: things[idx].color,
            })
        }
    }

    for p in model.projectiles.iter_mut() {
        let target = model.things[p.dest].position;
        p.position += (target - p.position).normalize() * 2.0;
    }
}

fn view(app: &App, model: &Model, frame: Frame) {
    let draw = app.draw();
    let time = app.elapsed_frames() as f32 / 60.0;

    if app.elapsed_frames() < 2 {
        draw.background().color(BLACK);
    }
    draw.rect()
        .w_h(1024.0, 1024.0)
        .color(srgba(0.0, 0.0, 0.0, 0.05));

    for thing in model.things.iter() {
        let radius = (time - thing.last_hit).min(2.0) * 2.0;
        draw.ellipse()
            .xy(thing.position)
            .color(thing.color)
            .radius(radius);
    }
    for proj in model.projectiles.iter() {
        draw.ellipse()
            .xy(proj.position)
            .color(proj.color)
            .radius(1.0);
    }

    draw.to_frame(app, &frame).unwrap();
}
