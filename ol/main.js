import './style.css';
import 'ol/ol.css';
import ImageLayer from 'ol/layer/Image';
import Map from 'ol/Map';
import Projection from 'ol/proj/Projection';
import Static from 'ol/source/ImageStatic';
import View from 'ol/View';
import {getCenter} from 'ol/extent';

const extent = [0, 0, 3000, 3000];
const projection = new Projection({
  code: 'aoc-image',
  units: 'pixels',
  extent: extent,
});

const map = new Map({
  layers: [
    new ImageLayer({
      source: new Static({
        url: 'http://localhost:1234/getImage',
        projection: projection,
        imageExtent: extent,
      }),
    }),
  ],
  target: 'map',
  view: new View({
    projection: projection,
    center: getCenter(extent),
    zoom: 1.8,
    maxZoom: 8,
  }),
});