export function sortPlayers(players, field, order) {
  return [...players].sort((a, b) => {
    if (order === 'asc') {
      return a[field] - b[field]
    }
    return b[field] - a[field]
  })
}