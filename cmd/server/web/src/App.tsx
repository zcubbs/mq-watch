import QueryPage from "@/pages/mq-watch/page.tsx";
import {ThemeProvider} from "@/context/theme-context";
import Header from "@/components/header.tsx";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";

// Initialize a new query client
const queryClient = new QueryClient();

function App() {

  return (
    <QueryClientProvider client={queryClient}>
      <ThemeProvider>
        <div className="app">
          <Header/>
          <QueryPage/>
        </div>
      </ThemeProvider>
    </QueryClientProvider>
  )
}

export default App
