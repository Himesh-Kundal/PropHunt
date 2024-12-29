import { getPlayerRank } from '../utils/ranks'

export function RankBadge({ wins }) {
  const rank = getPlayerRank(wins)
  
  return (
    <span className={`inline-flex items-center px-3 py-1 rounded bg-gray-800 ${rank.color} border border-opacity-20 border-current`}>
      <span className="mr-1">{rank.icon}</span>
      {rank.name}
    </span>
  )
}