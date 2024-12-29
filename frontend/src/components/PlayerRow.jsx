import { RankBadge } from './RankBadge'
import { calculateKDRatio, calculateWinRate } from '../utils/ranks'

export function PlayerRow({ player }) {
  const kdRatio = calculateKDRatio(player.kills, player.deaths)
  const winRate = calculateWinRate(player.wins, player.losses)

  return (
    <tr className="border-b border-gray-800 hover:bg-gray-800/50 transition-colors">
      <td className="p-3">
        <div className="flex items-center space-x-3">
          <span className="font-medium text-gray-100">{player.username}</span>
          <RankBadge wins={player.wins} />
        </div>
      </td>
      <td className="p-3">
        <div className="flex items-center">
          <span className="font-medium text-red-500">{player.kills}</span>
          <span className="text-xs text-gray-500 ml-2">({kdRatio} K/D)</span>
        </div>
      </td>
      <td className="p-3 text-gray-400">{player.deaths}</td>
      <td className="p-3">
        <div className="flex items-center">
          <span className="font-medium text-emerald-500">{player.wins}</span>
          <span className="text-xs text-gray-500 ml-2">({winRate}%)</span>
        </div>
      </td>
      <td className="p-3 text-gray-400">{player.losses}</td>
      <td className="p-3 text-gray-400">{player.draws}</td>
      <td className="p-3 text-amber-400">
        {Math.floor(player.time_alive / 60)}m {player.time_alive % 60}s
      </td>
    </tr>
  )
}