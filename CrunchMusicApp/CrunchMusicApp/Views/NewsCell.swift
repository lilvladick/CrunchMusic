import SwiftUI

struct NewsCell: View {
    let news: News
    
    var formatter: Formatter {
        Formatter(timestamp: news.publishedAt )
    }
    
    var body: some View {
        HStack {
            VStack(alignment: .leading) {
                Text(news.title.prefix(20) + (news.newsContent.count > 20 ? "..." : ""))
                    .font(.headline)
                Text(news.newsContent.prefix(20) + (news.newsContent.count > 20 ? "..." : "")).foregroundStyle(.secondary).font(.caption)
            }
            
            Spacer()
            
            VStack(alignment: .trailing) {
                Text(formatter.getFormattedDate())
                Text(formatter.getFormattedTime())
                Text(news.isBreaking ? "Breaking" : "").font(.caption2).foregroundStyle(.red).bold()
            }
            .foregroundStyle(.secondary)
            .font(.caption)
        }.bold()
    }
}

