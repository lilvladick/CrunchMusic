import Foundation

@MainActor
final class BreakingNewViewModel: ObservableObject {
    @Published var newsList: [News] = []
    @Published var isLoading: Bool = false
    @Published var error: String?
    
    private let newsURL = "http://127.0.0.1:8080/news"
    
    func loadBreakingNews() async {
        do {
            let news: [News] = try await NetworkManager.fetchData(from: newsURL)
            let breaking = news.filter { news in
                news.isBreaking
            }
            self.newsList = breaking
            self.isLoading = false
        } catch {
            self.error = "Error loading data: \(error.localizedDescription)"
            self.isLoading = false
        }
    }
}
