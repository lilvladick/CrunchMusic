import SwiftUI

struct MainScreenView: View {
    @StateObject var mainScreenViewModel: MainScreenViewModel
    @Binding var isTabBarHidden: Bool
    
    var body: some View {
        NavigationStack {
            VStack {
                if mainScreenViewModel.isLoading {
                    ProgressView("Загрузка новостей...")
                        .progressViewStyle(CircularProgressViewStyle())
                        .padding()
                } else if let error = mainScreenViewModel.error {
                    Text(error)
                } else {
                    List(mainScreenViewModel.newsList) { newsItem in
                        NavigationLink(destination: NewsDetailsView(newsCommentsViewModel: NewsCommentsViewModel(newsID: newsItem.id), news: newsItem)) {
                            NewsCell(news: newsItem)
                        }
                    }
                }
            }
            .refreshable {
                await mainScreenViewModel.loadNews()
            }
            .onAppear {
                Task {
                    await mainScreenViewModel.loadNews()
                }
            }
            .navigationTitle("Новости")
        }
    }
}

