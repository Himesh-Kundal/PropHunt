export function SortButton({ field, label, currentSort, onSort }) {
  const isActive = currentSort.field === field
  
  return (
    <th className="p-3">
      <button
        onClick={() => onSort(field)}
        className={`w-full text-left font-bold hover:text-red-500 transition-colors ${
          isActive ? 'text-red-500' : 'text-gray-400'
        }`}
      >
        {label} {isActive && (currentSort.order === 'asc' ? '↑' : '↓')}
      </button>
    </th>
  )
}