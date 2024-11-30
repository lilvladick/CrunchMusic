import Foundation

@MainActor
final class MainScreenViewModel: ObservableObject {
    @Published var newsList: [News] = []
    @Published var isLoading: Bool = false
    @Published var error: String?
    
    private let newsURL = "http://127.0.0.1:8080/news"
    
    func loadNews() async {
        do {
            let news: [News] = try await NetworkManager.fetchData(from: newsURL)
            self.newsList = news
            self.isLoading = false
        } catch {
            self.error = "Error loading data: \(error.localizedDescription)"
            self.isLoading = false
        }
    }
}
