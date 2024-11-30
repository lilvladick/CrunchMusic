import Foundation

@MainActor
final class NewsCommentsViewModel: ObservableObject {
    @Published var commentsList: [Comments] = []
    @Published var isLoading: Bool = true 
    @Published var error: String?
    
    private let commentsURL = "http://127.0.0.1:8080/commentsbynewsid?news_id="
    let newsID: Int
    
    init(newsID: Int) {
        self.newsID = newsID
    }
    
    func loadNewsCommentsData() async {
        let url = commentsURL + String(newsID)
        
        do {
            let comments: [Comments] = try await NetworkManager.fetchData(from: url)
            self.commentsList = comments
            self.isLoading = false
        } catch {
            self.error = "Error loading data: \(error.localizedDescription)"
            self.isLoading = false
        }
    }
}

