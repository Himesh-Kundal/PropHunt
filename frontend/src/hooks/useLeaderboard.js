import { useState, useEffect } from 'react'
import { userService } from '../services/api'
import { sortPlayers } from '../utils/sorting'

export function useLeaderboard() {
  const [players, setPlayers] = useState([])
  const [sortField, setSortField] = useState('Wins')
  const [sortOrder, setSortOrder] = useState('desc')
  const [topPlayer, setTopPlayer] = useState(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  useEffect(() => {
    const fetchPlayers = async () => {
      try {
        const data = await userService.getAllUsers()
        // Transform the data to match our frontend structure
        const transformedPlayers = data.map(player => ({
          username: player.Username,
          kills: player.Kills,
          deaths: player.Deaths,
          wins: player.Wins,
          losses: player.Losses,
          draws: player.Draws,
          time_alive: player.TimeAlive
        }))
        
        const sorted = sortPlayers(transformedPlayers, sortField.toLowerCase(), sortOrder)
        setPlayers(sorted)

        // Set top player based on most wins
        const top = [...transformedPlayers].sort((a, b) => b.wins - a.wins)[0]
        setTopPlayer(top)
        setLoading(false)
      } catch (err) {
        setError("Failed to load leaderboard data. \n Make sure You are connnected to Institute Network.")
        setLoading(false)
      }
    }

    fetchPlayers()
  }, [sortField, sortOrder])

  const handleSort = (field) => {
    if (field === sortField) {
      setSortOrder(sortOrder === 'asc' ? 'desc' : 'asc')
    } else {
      setSortField(field)
      setSortOrder('desc')
    }
  }

  return {
    players,
    sortField,
    sortOrder,
    handleSort,
    topPlayer,
    loading,
    error
  }
}