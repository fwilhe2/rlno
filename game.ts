import * as fs from 'fs';

interface League {
    Name: String
    Teams: Array<Team>
    Matches: Array<MatchData>
}

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
    homeTeam: string
    homeGoals: number
    awayTeam: string
    awayGoals: number
    events: Array<string>
}

interface MatchData {
    homeTeam: Team
    awayTeam: Team
}

const teams: Array<Team> = JSON.parse(fs.readFileSync("data.json").toString())

const game: MatchData = {
    homeTeam: teams[0],
    awayTeam: teams[9]
}

class Match {
    homeTeam: Team
    awayTeam: Team

    constructor(h: Team, a: Team) {
        this.homeTeam = h
        this.awayTeam = a
    }
}

function calculateMatchesForSeason(teams: Array<Team>): Array<Match> {
    return teams.flatMap(t => {
        return teams.filter(tt => tt != t).map(ttt => new Match(t, ttt))

    })

}

function playMatch(game: Match): GameResult {
    console.log(`Game ${game.homeTeam.Name} vs ${game.awayTeam.Name}`)
    let minute = 0
    let result: GameResult = {
        homeGoals: 0,
        homeTeam: game.homeTeam.Name,
        awayGoals: 0,
        awayTeam: game.awayTeam.Name,
        events: []
    }

    do {
        if (Math.random() < 0.01) {
            result.homeGoals++
            console.log(`[${minute}]: Goal for ${game.homeTeam.Name}, standings: ${result.homeGoals}:${result.awayGoals}`)
        }
        if (Math.random() < 0.01) {
            result.awayGoals++
            console.log(`[${minute}]: Goal for ${game.awayTeam.Name}, standings: ${result.homeGoals}:${result.awayGoals}`)
        }

        minute++
    } while (minute < 90)

    return result
}

// console.log(playMatch(game))

const matches = calculateMatchesForSeason(teams)
const results = matches.map(m => playMatch(m))
fs.writeFileSync('match-results.json', JSON.stringify(results, null, 4))

interface Standing {

    team: string
    points: Number
}

function calculateStandings(matchResults: Array<GameResult>): Array<Standing> {
    return teams.map(t => { return { team: t.Name, points: calculateStandingsForTeam(t.Name, matchResults) } }).sort((a, b) => b.points - a.points)
}

function calculateStandingsForTeam(team: string, matchResults: Array<GameResult>) {
    return matchResults.filter(m => m.homeTeam == team).filter(m => m.homeGoals > m.awayGoals).length * 3 +
        matchResults.filter(m => m.awayTeam == team).filter(m => m.awayGoals > m.homeGoals).length * 3 +
        matchResults.filter(m => m.homeTeam == team || m.awayTeam == team).filter(m => m.homeGoals === m.awayGoals).length * 1
}

const standings = calculateStandings(results)
fs.writeFileSync('standings.json', JSON.stringify(standings, null, 4))
