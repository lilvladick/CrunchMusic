import SwiftUI

struct NewsDetailsView: View {
    @ObservedObject var newsCommentsViewModel: NewsCommentsViewModel
    
    let news: News
    var formatter: Formatter {
        Formatter(timestamp: news.publishedAt )
    }
    
    var body: some View {
        VStack(alignment: .leading) {
            VStack{
                HStack {
                    Text(news.title)
                        .font(.title)
                        .bold()
                }
                
                HStack(alignment: .lastTextBaseline) {
                    Text(news.isBreaking ? "Breaking" : "")
                        .foregroundStyle(.red)
                    Spacer()
                    
                    Text(formatter.getFormattedDate())
                    Text(formatter.getFormattedTime())
                }
                .foregroundStyle(Color.gray)
                .font(.caption)
            }
            .padding()
            .background(Color.gray.opacity(0.1))
           
           VStack(alignment: .leading) {
               Text(news.newsContent)
           }
           .padding(10)
            
            CommentsView(newsCommentsViewModel:  newsCommentsViewModel)
        }
       .onAppear{
           Task{
               await newsCommentsViewModel.loadNewsCommentsData()
           }
       }
    }
}
