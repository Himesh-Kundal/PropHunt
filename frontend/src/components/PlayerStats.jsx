export function PlayerStats({ stat, label, className = '' }) {
  return (
    <div className={`text-center ${className}`}>
      <div className="text-2xl font-bold text-gray-100">{stat}</div>
      <div className="text-sm text-gray-400 uppercase tracking-wider">{label}</div>
    </div>
  )
}