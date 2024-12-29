// Military-inspired ranks with weapon tiers
export const ranks = [
  { name: 'Recruit', minWins: 0, color: 'text-gray-400', icon: '🔰' },
  { name: 'Gunner', minWins: 20, color: 'text-red-500', icon: '🔫' },
  { name: 'Sniper', minWins: 40, color: 'text-emerald-500', icon: '🎯' },
  { name: 'Commando', minWins: 60, color: 'text-purple-500', icon: '💥' },
  { name: 'Warlord', minWins: 80, color: 'text-amber-500', icon: '⚔️' },
  { name: 'Predator', minWins: 95, color: 'text-red-600', icon: '☠️' }
]

export function getPlayerRank(wins) {
  return ranks.reduce((highest, rank) => 
    wins >= rank.minWins ? rank : highest
  )
}

export function calculateKDRatio(kills, deaths) {
  return deaths === 0 ? kills : (kills / deaths).toFixed(2)
}

export function calculateWinRate(wins, losses) {
  const total = wins + losses
  return total === 0 ? 0 : ((wins / total) * 100).toFixed(1)
}