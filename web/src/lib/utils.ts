import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

export const API_URL = import.meta.env.VITE_API_URL;

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export const getRandomColor = (): string => {
  const letters = '0123456789ABCDEF';
  let color = '#';
  for (let i = 0; i < 6; i++) {
    color += letters[Math.floor(Math.random() * 16)];
  }
  return color;
};

export const getRandomPastelColor = (): string => {
  const mix = [255, 255, 255]; // Mix the color with white to ensure pastel
  let red = Math.floor(Math.random() * 256);
  let green = Math.floor(Math.random() * 256);
  let blue = Math.floor(Math.random() * 256);

  // Mix the color with white
  red = Math.floor((red + mix[0]) / 2);
  green = Math.floor((green + mix[1]) / 2);
  blue = Math.floor((blue + mix[2]) / 2);

  const redHex = red.toString(16).padStart(2, '0');
  const greenHex = green.toString(16).padStart(2, '0');
  const blueHex = blue.toString(16).padStart(2, '0');

  return `#${redHex}${greenHex}${blueHex}`;
};


// Function to format date to RFC3339
export function formatDateToRFC3339(date: Date) {
  function pad(value: number) {
    return value < 10 ? `0${value}` : value;
  }

  return `${date.getUTCFullYear()}-${pad(date.getUTCMonth() + 1)}-${pad(date.getUTCDate())}T${pad(date.getUTCHours())}:${pad(date.getUTCMinutes())}:${pad(date.getUTCSeconds())}Z`;
}
