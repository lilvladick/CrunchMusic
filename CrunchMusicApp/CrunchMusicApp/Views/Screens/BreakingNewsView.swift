import SwiftUI

struct BreakingNewsView: View {
    @StateObject var breakingNewsViewModel: BreakingNewViewModel
    @Binding var isTabBarHidden: Bool
    
    var body: some View {
        NavigationStack {
            VStack {
                if breakingNewsViewModel.isLoading {
                    ProgressView("Загрузка новостей...")
                        .progressViewStyle(CircularProgressViewStyle())
                        .padding()
                } else if let error = breakingNewsViewModel.error {
                    Text(error)
                } else {
                    List(breakingNewsViewModel.newsList) { newsItem in
                        NavigationLink(destination: NewsDetailsView(newsCommentsViewModel: NewsCommentsViewModel(newsID: newsItem.id), news: newsItem)) {
                            NewsCell(news: newsItem)
                        }
                    }
                }
            }
            .refreshable {
                await breakingNewsViewModel.loadBreakingNews()
            }
            .onAppear {
                Task {
                    await breakingNewsViewModel.loadBreakingNews()
                }
            }
            .navigationTitle("Срочные новости")
        }
    }
}
