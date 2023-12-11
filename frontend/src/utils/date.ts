export function formatDate(date: Date): string {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}

export function getDateRange(intervalDays: number): [string, string] {
  const currentDate = new Date();
  const formattedCurrentDate = formatDate(currentDate);
  
  const pastDate = new Date();
  pastDate.setDate(pastDate.getDate() - intervalDays);
  const formattedPastDate = formatDate(pastDate);
  
  return [formattedCurrentDate, formattedPastDate];
}