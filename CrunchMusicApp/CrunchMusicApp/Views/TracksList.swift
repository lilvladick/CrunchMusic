import SwiftUI

struct TracksList: View {
    @State private var isShowingMenu:Bool = false
    var body: some View {
        ZStack {
            NavigationStack() {
                VStack {
                    // JSON PARSE TRACKS TO LIST
                }
                .toolbar {
                    ToolbarItem(placement: .topBarLeading, content: {
                        Button(action: {
                            withAnimation {
                                isShowingMenu.toggle()
                            }
                        }, label: {
                            Image(systemName: "list.dash")
                        })
                    })
                    
                    ToolbarItem(placement: .topBarTrailing, content: {
                        Button(action: {
                            // Serch action
                        }, label: {
                            Image(systemName: "magnifyingglass")
                        })
                    })
                }
                .bold()
            }
            if isShowingMenu {
                Color.black.opacity(0.3)
                    .ignoresSafeArea(.all)
                    .onTapGesture {
                        withAnimation {
                            isShowingMenu = false
                        }
                    }
            }
            
            SideMenuView(isShowing: $isShowingMenu)
                .transition(.move(edge: .leading))
        }
    }
}

#Preview {
    TracksList()
}
