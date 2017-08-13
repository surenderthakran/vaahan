'use strict';

var Utils = {};

Utils.getX = function(point) {
  return point.x;
}

Utils.getY = function(point) {
  return _trackData.height - point.y;
}
