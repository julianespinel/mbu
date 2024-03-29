# MBU [![Build Status](https://travis-ci.org/julianespinel/mbu.svg?branch=master)](https://travis-ci.org/julianespinel/mbu)
MBU (Machine Basic Usage) is a system that provides basic machine information via HTTP.

With MBU you can get the following information from the machine it is installed on.

* CPU usage
* RAM usage
* Disk usage

## Installation

MBU uses the port 7007 by default. If you want to change it just change the port number in the file: [mbu.toml](https://github.com/julianespinel/mbu/blob/master/mbu.toml).

### Linux

Execute these steps in you terminal.

1. `git clone git@github.com:julianespinel/mbu.git`
2. `cd mbu`
3. `./mbu`

## Usage

You can access the information provided by MBU directly from your web browser or from an HTTP client.

You can copy and paste the following lines in your browser to see the information from the host MBU is installed on: 

1. `http://localhost:7007/mbu/admin/ping`
2. `http://localhost:7007/mbu/api/cpu`
3. `http://localhost:7007/mbu/api/ram`
4. `http://localhost:7007/mbu/api/disk`
5. `http://localhost:7007/mbu/api/all`

For more information about the services offered by MBU please check the [API](https://github.com/julianespinel/mbu#api) section of this readme.

### Example

*Request:*

`GET http://localhost:7007/mbu/api/all`

*Response:*

```json
{
    "cpu": {
        "numberOfCores": 4,
        "averageUsagePercentage": 10.801282,
        "usagePercentagePerCore": [
            5,
            23.076923,
            5.1282053,
            10
        ]
    },
    "ram": {
        "totalGB": 7.612256,
        "usedGB": 2.0867763,
        "availableGB": 5.52548,
        "usagePercentage": 27.413372
    },
    "disk": {
        "totalGB": 976.49677,
        "usedGB": 53.40926,
        "availableGB": 923.08746,
        "usagePercentage": 5.4694767
    }
}
```

## API

MBU provides the following 5 services: 

### 1. Ping

Use the ping service to check the health of MBU.

*Request:*

`GET http://localhost:7007/mbu/admin/ping`

*Response:*

`pong`

### 2. CPU

Provides CPU usage information.

*Request:*

`GET http://localhost:7007/mbu/api/cpu`

*Response:*

```json
{
    "numberOfCores": 4,
    "averageUsagePercentage": 13.351047,
    "usagePercentagePerCore": [
        7.692308,
        22.5,
        5.263158,
        17.948719
    ]
}
```

### 3. RAM

Provides RAM usage information.

*Request:*

`GET http://localhost:7007/mbu/api/ram`

*Response:*

```json
{
    "totalGB": 7.612256,
    "usedGB": 2.9295201,
    "availableGB": 4.682736,
    "usagePercentage": 38.484257
}
```

### 4. Disk

Provides disk usage information.

*Request:*

`GET http://localhost:7007/mbu/api/disk`

*Response:*

```json
{
    "totalGB": 976.49677,
    "usedGB": 53.44063,
    "availableGB": 923.0561,
    "usagePercentage": 5.472689
}
```

### 5. All

Provides CPU, RAM and disk information in one JSON.

*Request:*

`GET http://localhost:7007/mbu/api/all`

*Response:*

```json
{
    "cpu": {
        "numberOfCores": 4,
        "averageUsagePercentage": 10.801282,
        "usagePercentagePerCore": [
            5,
            23.076923,
            5.1282053,
            10
        ]
    },
    "ram": {
        "totalGB": 7.612256,
        "usedGB": 2.0867763,
        "availableGB": 5.52548,
        "usagePercentage": 27.413372
    },
    "disk": {
        "totalGB": 976.49677,
        "usedGB": 53.40926,
        "availableGB": 923.08746,
        "usagePercentage": 5.4694767
    }
}
```

## License

This project is under the MIT License. See the [LICENSE](https://github.com/julianespinel/mbu/blob/master/LICENSE) file for the full license text.
