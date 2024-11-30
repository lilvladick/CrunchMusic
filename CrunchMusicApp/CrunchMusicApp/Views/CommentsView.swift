import SwiftUI

struct CommentsView: View {
    @ObservedObject var newsCommentsViewModel: NewsCommentsViewModel
    
    var body: some View {
        VStack {
            HStack{
                Text("Комментарии")
                    .font(.title2)
                    .bold()
                Spacer()
                Button(action: {
                    //
                }, label: {
                    Image(systemName: "plus")
                })
            }
            .padding(.horizontal)
            
            List(newsCommentsViewModel.commentsList) { commentItem in
                CommentCell(comment: commentItem)
            }
            .listStyle(.plain)
        }
    }
}
