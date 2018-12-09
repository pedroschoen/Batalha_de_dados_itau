from datetime import datetime

import json
import flask
import numpy as np
import pandas as pd
from dateutil.relativedelta import relativedelta
from flask import request
from scipy.spatial import distance

# initialize our Flask application
app = flask.Flask(__name__)


def get_best_route(base, starting_point):
    geo_points = [np.array(starting_point)]
    lat_long_cols = ['LATITUDE', 'LONGITUDE']

    for i in range(5):
        base['dist'] = base.apply(
            lambda x: distance.euclidean(x[lat_long_cols], geo_points[i - 1]), axis=1)
        sorted_vals = base.sort_values('dist')
        for j in range(3):
            next_point = sorted_vals.iloc[j][lat_long_cols].values
            if (next_point == geo_points).any():
                continue
            geo_points.append(next_point)

        if ((next_point == geo_points[0]).all()) or (next_point == geo_points[1]).all():
            next_point = geo_points[0]
            geo_points.append(next_point)
            break

    c = 0
    while (next_point != geo_points[0]).all():
        mean_point = (geo_points[0] + geo_points[-1]) / 2
        base['dist'] = base.apply(
            lambda x: distance.euclidean(x[lat_long_cols], mean_point), axis=1)
        sorted_vals = base.sort_values('dist')
        next_point = sorted_vals.iloc[c][lat_long_cols].values

        if (next_point != geo_points[0]).all() and (next_point != geo_points[1]).all() and not (
                next_point == geo_points).any():
            geo_points.append(next_point)

        c += 1
        if c > 5:
            geo_points.append(geo_points[0])
            break
    return geo_points


@app.route('/', methods=['GET'])
def index():
    return 'ok'


@app.route('/recommend_route', methods=['POST'])
def recommend_route():
    # geo point
    lat = request.form['lat']
    long = request.form['long']
    starting_point = [float(lat), float(long)]

    ref_date = datetime.now() - relativedelta(month=1)
    rel_time1 = datetime.now() - relativedelta(hours=1)
    rel_time2 = datetime.now() + relativedelta(hours=1)

    time_col = 'BO_INICIADO'
    df = pd.read_csv('clean.csv', parse_dates=[time_col])
    df = df[(df[time_col] <= ref_date) &
            (rel_time1.time() < df[time_col].dt.time) &
            (df[time_col].dt.time < rel_time2.time())]

    points = get_best_route(df, starting_point)

    # return json.dumps(points)
    data = []
    for p in points:
        data.append(p.tolist())

    return flask.jsonify({'points': data})


if __name__ == '__main__':
    app.run(debug=True)
