import SwiftUI

struct MainScreenView: View {
    @State private var newsList: [News] = []
    @State private var isLoading: Bool = true
    @State private var error: String?
    
    let newsURL = "http://127.0.0.1:8080/news"
    
    var body: some View {
        NavigationStack {
            VStack {
                if isLoading {
                    ProgressView("Загрузка новостей...")
                        .progressViewStyle(CircularProgressViewStyle())
                        .padding()
                } else if let error = error {
                    Text(error)
                } else {
                    List(newsList) { newsItem in
                        NavigationLink(destination: Text(newsItem.newsContent)) {
                            NewsCell(news: newsItem)
                        }
                    }
                }
            }.onAppear {
                Task {
                    await loadNews()
                }
            }
            .navigationTitle("Новости")
        }
    }
    
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

#Preview {
    MainScreenView()
}
