import SwiftUI

struct CommentCell: View {
    let comment: Comments
    var formatter: Formatter {
        Formatter(timestamp: comment.createdAt)
    }
    
    var body: some View {
        VStack {
            HStack(alignment: .top) {
                VStack(alignment: .center) {
                    Image(systemName: "person.fill")
                        .font(.title3)
                        .padding(10)
                        .background(Color.black)
                        .clipShape(.circle)
                        .foregroundStyle(Color.white)
                    Text("user")
                        .bold()
                        .font(.caption)
                }
                
                Text(comment.commentContent)
                    .padding(.horizontal,5)
                    .font(.callout)
                
                Spacer()
                
                VStack(alignment: .trailing) {
                    Text(formatter.getFormattedDate())
                    Text(formatter.getFormattedTime())
                }
                .font(.caption)
                .bold()
                .foregroundStyle(Color.gray)
            }
        }
    }
}

#Preview {
    CommentCell(comment: Comments(id: 1, newsId: 1, authorId: 1, commentContent: "Ужас какой", createdAt: "2024-11-29T15:51:39.227498Z", updatedAt: "2024-11-29T15:51:39.227498Z"))
}
