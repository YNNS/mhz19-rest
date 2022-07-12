<h1 align="center">MH-Z19 CO₂ Rest</h1>

<p align="center"> 
  <img src="https://img.shields.io/badge/made%20by-YNNS-blueviolet.svg?style=for-the-badge" >

  <img src="https://img.shields.io/badge/go-1.18-green.svg?style=for-the-badge">

  <img src="https://img.shields.io/github/stars/silent-lad/Vue2BaremetricsCalendar.svg?style=for-the-badge">

  <img src="https://img.shields.io/github/languages/top/silent-lad/Vue2BaremetricsCalendar.svg?style=for-the-badge">

  <img src="https://img.shields.io/github/issues/silent-lad/Vue2BaremetricsCalendar.svg?style=for-the-badge">

  <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge">
</p>

_A simple rest server running on your raspberry pi, reporting the current CO2 concentration in your room._

---

This project is part of the blog
article [How I built my own CO2 monitor using a Raspberry Pi for Home Office](https://medium.com/@YNNSme/how-i-built-my-own-co2-monitor-using-a-raspberry-pi-for-home-office-c8b168fcadfa).

![Project showcase](https://user-images.githubusercontent.com/29334025/178571309-c5ecb4dc-14e8-4f87-9750-eb1b46d5c289.png)

# Installation

### 1. Download the pre-built binary

* Go to the [release page]() and download the latest binary for your architecture.
* Add the execute permission to the binary. This is required to run the binary.
  `chmod +x ./mhz19-rest`
* Execute the binary on your Raspberry Pi.
  `./mhz19-rest`
* You can now access the rest api on `http://[raspberry-ip]:3500/api/v1/co2`

### 2. Build it yourself

* Clone this repository. `git clone git@github.com:YNNS/mhz19-rest.git`
  or `git clone https://github.com/YNNS/mhz19-rest.git`
* Build your binary
    * on the remote machine: `go build -o bin/mhz19-rest main.go`
    * on a local machine: `GOOS=linux GOARCH=arm64 go build -o bin/mhz19-rest main.go` (or `GOARCH=arm` if you are using
      a 32-bit ARM image)
* `cd bin/` - Or upload the file to the remote machine.
* Add the execute permission to the binary. This is required to run the binary.
  `chmod +x ./mhz19-rest`
* Execute the binary on your Raspberry Pi.
  `./mhz19-rest`
* You can now access the rest api on `http://[raspberry-ip]:3500/api/v1/co2`

## Frontend

Check out the corresponding [frontend here](https://github.com/YNNS/mhz19-react-ui)



## 👍 Contribute

If you want to say **thank you**

1. Add a [GitHub Star](https://github.com/YNNS/mhz19-rest/stargazers) to the project.
2. Tweet about the project [on your Twitter](https://twitter.com/intent/tweet?text=%E2%9D%A4%EF%B8%8F%E2%80%8D%F0%9F%94%A5%20Check%20out%20how%20to%20build%20your%20own%20CO%E2%82%82%20monitor%3A%20https%3A%2F%2Fgithub.com%2FYNNS%2Fmhz19-rest).
3. Add some claps to the story of this project on [Medium](https://medium.com/@YNNSme/how-i-built-my-own-co2-monitor-using-a-raspberry-pi-for-home-office-c8b168fcadfa)
