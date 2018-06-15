$(document).ready(function() {


    var moi = [];
    var totalE = [];
    var totalS = [];

    var mois = [];
    var prix = [];
    var ctx2 = document.getElementById("myAreaChart");
    var ctx3 = document.getElementById("myPieChart");

    $.getJSON('/json', function (data) {
        $.each(data, function (entryIndex, entry) {
            mois.push(entry['mois']);
            prix.push(entry['prix']);
        });
// -- Bar Chart Example
        var ctx = document.getElementById("myBarChart");
        var myLineChart = new Chart(ctx, {
            type: 'bar',
            data: {
                labels: mois,
                datasets: [{
                    label: "Quantité",
                    backgroundColor: "rgba(2,117,216,1)",
                    borderColor: "rgba(2,117,216,1)",
                    data: prix,
                }],
            },
            options: {
                scales: {
                    xAxes: [{
                        time: {
                            unit: 'month'
                        },
                        gridLines: {
                            display: false
                        },
                        ticks: {
                            maxTicksLimit: 6
                        }
                    }],
                    yAxes: [{
                        ticks: {
                            interval: 2,
                            min: 0,
                            max: Math.max(...prix),
                            maxTicksLimit: Math.floor(Math.max(...prix)/5)
                        },
                        gridLines: {
                            display: true
                        }
                    }],
                },
                legend: {
                    display: false
                }
            }
        });

    })

///////////////////EPHRAU'S CHART

        $.getJSON('/ephraujson', function (data) {
            $.each(data, function (entryIndex, entry) {
                moi.push(entry['moi']);
                totalE.push(entry['totalE']);
                totalS.push(entry['totalS']);
            });
            var ctx = document.getElementById("mymy");
            var myLineChart = new Chart(ctx, {
                type: 'bar',
                data: {
                    labels: moi,
                    datasets: [{
                        label: "Entrée",
                        backgroundColor: "rgba(140, 140, 140,0.8)",
                        borderColor: "rgba(140, 140, 140,1)",
                        data: totalE
                    }, {
                        label: "Sortie",
                        backgroundColor: "rgba(255,153,0,0.8)",
                        borderColor: "rgba(255,153,0,1)",
                        data: totalS
                    }],
                },
                options: {
                    scales: {
                        xAxes: [{
                            time: {
                                unit: 'month'
                            },
                            gridLines: {
                                display: true
                            },
                            ticks: {
                                maxTicksLimit: 12
                            }
                        }],
                        yAxes: [{
                            ticks: {
                                min: 0,
                                max: Math.max(Math.max(...totalE), Math.max(...totalS)),
                                maxTicksLimit: Math.floor(Math.max(Math.max(...totalE), Math.max(...totalS))/5)
                            },
                            gridLines: {
                                display: true
                            }
                        }],
                    },
                    legend: {
                        display: false
                    }
                }
            });
            var myLineChart2 = new Chart(ctx2, {
                type: 'line',
                data: {
                    labels: moi,
                    datasets: [{
                        label: "Coût",
                        lineTension: 0.3,
                        backgroundColor: "rgba(255, 153, 0,0.2)",
                        borderColor: "rgba(255, 153, 0,1)",
                        pointRadius: 5,
                        pointBackgroundColor: "rgba(255, 153, 0,1)",
                        pointBorderColor: "rgba(255,255,255,0.8)",
                        pointHoverRadius: 5,
                        pointHoverBackgroundColor: "rgba(140, 140, 140,1)",
                        pointHitRadius: 20,
                        pointBorderWidth: 2,
                        data: totalE,
                    }],
                },
                options: {
                    scales: {
                        xAxes: [{
                            time: {
                                unit: 'date'
                            },
                            gridLines: {
                                display: false
                            },
                            ticks: {
                                maxTicksLimit: 12
                            }
                        }],
                        yAxes: [{
                            ticks: {
                                min: 0,
                                max: Math.max(...totalE),
                                maxTicksLimit: 9
                            },
                            gridLines: {
                                color: "rgba(0, 0, 0, .125)",
                            }
                        }],
                    },
                    legend: {
                        display: false
                    }
                }
            });

        });

    return false;
})






