import MQWatchPage from "@/pages/mq-watch/page.tsx";
import {ThemeProvider} from "@/context/theme-provider";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";

// Initialize a new query client
const queryClient = new QueryClient();

function App() {

  return (
    <ThemeProvider defaultTheme="dark" storageKey=" vite-ui-theme">
      <QueryClientProvider client={queryClient}>
        <MQWatchPage/>
      </QueryClientProvider>
    </ThemeProvider>
  )
}

export default App
