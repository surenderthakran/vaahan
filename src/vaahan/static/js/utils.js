'use strict';

var Utils = {};

Utils.getX = function(point) {
  return point.x;
}

Utils.getY = function(point) {
  return _map.track.height - point.y;
}

window.Utils = Utils;
