import * as fs from 'fs';

interface Team {
    Name: string
    Stats: Stats
    Players: Array<Player>
}

interface Player {
    Name: string
}

interface Stats {
    AverageMarketValue: number
    TotalMarketValue: number
}

interface GameResult {
    homeGoals: number
    awayGoals: number
}

const teams: Array<Team> = JSON.parse(fs.readFileSync("data.json").toString())

const teamHome = teams[0]
const teamAway = teams[9]

function playMatch() {
    let minute = 0
    let result: GameResult = {
        homeGoals: 0,
        awayGoals: 0
    }

    do {
        if (Math.random() < 0.01) {
            result.homeGoals++
            console.log(`[${minute}]: Goal for ${teamHome.Name}, standings: ${result.homeGoals}:${result.awayGoals}`)
        }
        if (Math.random() < 0.01) {
            result.awayGoals++
            console.log(`[${minute}]: Goal for ${teamAway.Name}, standings: ${result.homeGoals}:${result.awayGoals}`)
        }

        minute++
    } while (minute < 90)

    return result
}

console.log(playMatch())