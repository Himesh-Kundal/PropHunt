import { useLeaderboard } from '../hooks/useLeaderboard'
import { SortButton } from './SortButton'
import { PlayerRow } from './PlayerRow'
import { PlayerStats } from './PlayerStats'

function Leaderboard() {
  const { players, sortField, sortOrder, handleSort, topPlayer, loading, error } = useLeaderboard()

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-900 text-gray-100 flex items-center justify-center">
        <div className="text-xl text-red-500">Loading leaderboard...</div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="min-h-screen bg-gray-900 text-gray-100 flex items-center justify-center">
        <div className="text-xl text-red-500">{error}</div>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-gray-900 text-gray-100">
      <div className="container mx-auto p-4">
        <h2 className="text-3xl font-bold mb-6 text-center bg-clip-text text-transparent bg-gradient-to-r from-red-500 via-purple-500 to-emerald-500">
          PROP HUNT LEADERBOARD
        </h2>
        
        {topPlayer && (
          <div className="bg-gradient-to-r from-gray-800 to-gray-900 rounded-lg p-6 mb-8 border border-gray-700 shadow-lg relative overflow-hidden">
            <div className="absolute inset-0 bg-black/50"></div>
            <div className="relative z-10">
              <h3 className="text-xl font-bold mb-4 flex items-center">
                <span className="text-red-500 mr-2">☠️</span>
                APEX PREDATOR
              </h3>
              <div className="flex justify-between items-center">
                <div className="flex items-center space-x-4">
                  <span className="text-2xl font-bold text-red-500">{topPlayer.username}</span>
                </div>
                <div className="flex space-x-12">
                  <PlayerStats stat={topPlayer.kills} label="Eliminations" />
                  <PlayerStats stat={topPlayer.wins} label="Victories" />
                  <PlayerStats 
                    stat={`${Math.floor(topPlayer.time_alive / 60)}m`} 
                    label="Survival Time" 
                  />
                </div>
              </div>
            </div>
          </div>
        )}

        <div className="bg-gray-800 rounded-lg shadow-2xl overflow-hidden border border-gray-700">
          <div className="overflow-x-auto">
            <table className="min-w-full">
              <thead className="bg-gray-900/50">
                <tr>
                  <SortButton
                    field="username"
                    label="OPERATOR"
                    currentSort={{ field: sortField, order: sortOrder }}
                    onSort={handleSort}
                  />
                  <SortButton
                    field="kills"
                    label="KILLS"
                    currentSort={{ field: sortField, order: sortOrder }}
                    onSort={handleSort}
                  />
                  <SortButton
                    field="deaths"
                    label="DEATHS"
                    currentSort={{ field: sortField, order: sortOrder }}
                    onSort={handleSort}
                  />
                  <SortButton
                    field="wins"
                    label="WINS"
                    currentSort={{ field: sortField, order: sortOrder }}
                    onSort={handleSort}
                  />
                  <SortButton
                    field="losses"
                    label="LOSSES"
                    currentSort={{ field: sortField, order: sortOrder }}
                    onSort={handleSort}
                  />
                  <SortButton
                    field="draws"
                    label="DRAWS"
                    currentSort={{ field: sortField, order: sortOrder }}
                    onSort={handleSort}
                  />
                  <SortButton
                    field="time_alive"
                    label="SURVIVAL TIME"
                    currentSort={{ field: sortField, order: sortOrder }}
                    onSort={handleSort}
                  />
                </tr>
              </thead>
              <tbody>
                {players.map((player) => (
                  <PlayerRow key={player.username} player={player} />
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Leaderboard