'use strict';

function drawBoundary() {
  console.log("inside drawBoundary()");
  let boundary = _map.track.boundary;
  _map.context.fillStyle = '#8ae291';
  _map.context.fillRect(Utils.getX(boundary.shape.top_left), Utils.getY(boundary.shape.top_left), boundary.shape.width, boundary.shape.height);
}

function drawCar() {
  console.log("inside drawCar()");
  _car.context.clearRect(0, 0, _car.canvas.width, _car.canvas.height)

  _car.context.lineWidth = 0.5;

  // draw car's border.
  _car.context.beginPath();
  _car.context.moveTo(Utils.getX(_car.data.points.front_left), Utils.getY(_car.data.points.front_left));
  _car.context.lineTo(Utils.getX(_car.data.points.front_right), Utils.getY(_car.data.points.front_right));
  _car.context.lineTo(Utils.getX(_car.data.points.back_right), Utils.getY(_car.data.points.back_right));
  _car.context.lineTo(Utils.getX(_car.data.points.back_left), Utils.getY(_car.data.points.back_left));
  _car.context.lineTo(Utils.getX(_car.data.points.front_left), Utils.getY(_car.data.points.front_left));
  _car.context.closePath();
  _car.context.stroke();
  _car.context.fillStyle = "yellow";
  _car.context.fill();

  // draw color points at car's corners.
  _car.context.fillStyle = "red";
  _car.context.fillRect(Utils.getX(_car.data.points.front_left) - 2, Utils.getY(_car.data.points.front_left) - 2, 4, 4);
  _car.context.fillRect(Utils.getX(_car.data.points.back_left) - 2, Utils.getY(_car.data.points.back_left) - 2, 4, 4);

  _car.context.fillStyle = "blue";
  _car.context.fillRect(Utils.getX(_car.data.points.front_right) - 2, Utils.getY(_car.data.points.front_right) - 2, 4, 4);
  _car.context.fillRect(Utils.getX(_car.data.points.back_right) - 2, Utils.getY(_car.data.points.back_right) - 2, 4, 4);

  // draw car's sensors vision.
  _car.context.lineWidth = 0.1;
  for (var i in _car.data.sensors) {
    if (_car.data.sensors[i].intersection !== null) {
      _car.context.moveTo(Utils.getX(_car.data.sensors[i].ray.start_point), Utils.getY(_car.data.sensors[i].ray.start_point));
      _car.context.lineTo(Utils.getX(_car.data.sensors[i].intersection), Utils.getY(_car.data.sensors[i].intersection));
      _car.context.stroke();
    }
  }
}
