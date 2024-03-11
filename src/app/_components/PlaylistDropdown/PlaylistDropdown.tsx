export interface Playlist {
  id: string
  name: string
}

interface PlaylistDropdownProps {
  playlists: Playlist[]
  setPlaylist: (playlist: Playlist) => void
  playlist?: Playlist
  className?: string
  disabled: boolean
}
export default function PlaylistDropdown({
  playlists: playlistsProps,
  setPlaylist: setPlaylistProps,
  className: classNameProps,
  playlist: playlistProps,
  disabled: disabledProps,
}: PlaylistDropdownProps) {
  return (
    <div className={classNameProps}>
      <details className="dropdown w-full">
        <summary
          className={`btn m-1 w-full ${disabledProps && 'btn-disabled'}`}
        >
          {playlistProps?.name ?? 'Choose a playlist'}
        </summary>

        <ul className="menu dropdown-content z-[1] w-52 rounded-box bg-base-100 p-2 shadow">
          {playlistsProps.map((playlist) => {
            return (
              <li key={playlist.id}>
                <a onClick={() => setPlaylistProps(playlist)}>
                  {playlist.name}
                </a>
              </li>
            )
          })}
        </ul>
      </details>
    </div>
  )
}
