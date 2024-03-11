// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package ebpftracer

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type tracerGlobalConfigT struct{ PidNsId uint64 }

// loadTracer returns the embedded CollectionSpec for tracer.
func loadTracer() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_TracerBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load tracer: %w", err)
	}

	return spec, err
}

// loadTracerObjects loads tracer and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*tracerObjects
//	*tracerPrograms
//	*tracerMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadTracerObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadTracer()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// tracerSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type tracerSpecs struct {
	tracerProgramSpecs
	tracerMapSpecs
}

// tracerSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type tracerProgramSpecs struct {
	CgroupBpfRunFilterSkb                    *ebpf.ProgramSpec `ebpf:"cgroup_bpf_run_filter_skb"`
	CgroupMkdirSignal                        *ebpf.ProgramSpec `ebpf:"cgroup_mkdir_signal"`
	CgroupRmdirSignal                        *ebpf.ProgramSpec `ebpf:"cgroup_rmdir_signal"`
	CgroupSkbEgress                          *ebpf.ProgramSpec `ebpf:"cgroup_skb_egress"`
	CgroupSkbIngress                         *ebpf.ProgramSpec `ebpf:"cgroup_skb_ingress"`
	LkmSeekerKsetTail                        *ebpf.ProgramSpec `ebpf:"lkm_seeker_kset_tail"`
	LkmSeekerModTreeTail                     *ebpf.ProgramSpec `ebpf:"lkm_seeker_mod_tree_tail"`
	LkmSeekerNewModOnlyTail                  *ebpf.ProgramSpec `ebpf:"lkm_seeker_new_mod_only_tail"`
	LkmSeekerProcTail                        *ebpf.ProgramSpec `ebpf:"lkm_seeker_proc_tail"`
	OomMarkVictim                            *ebpf.ProgramSpec `ebpf:"oom_mark_victim"`
	SchedProcessExecEventSubmitTail          *ebpf.ProgramSpec `ebpf:"sched_process_exec_event_submit_tail"`
	SendBin                                  *ebpf.ProgramSpec `ebpf:"send_bin"`
	SendBinTp                                *ebpf.ProgramSpec `ebpf:"send_bin_tp"`
	SysDupExitTail                           *ebpf.ProgramSpec `ebpf:"sys_dup_exit_tail"`
	SysEnterInit                             *ebpf.ProgramSpec `ebpf:"sys_enter_init"`
	SysEnterSubmit                           *ebpf.ProgramSpec `ebpf:"sys_enter_submit"`
	SysExitInit                              *ebpf.ProgramSpec `ebpf:"sys_exit_init"`
	SysExitSubmit                            *ebpf.ProgramSpec `ebpf:"sys_exit_submit"`
	SyscallAccept4                           *ebpf.ProgramSpec `ebpf:"syscall__accept4"`
	SyscallExecve                            *ebpf.ProgramSpec `ebpf:"syscall__execve"`
	SyscallExecveat                          *ebpf.ProgramSpec `ebpf:"syscall__execveat"`
	SyscallInitModule                        *ebpf.ProgramSpec `ebpf:"syscall__init_module"`
	TraceRegisterChrdev                      *ebpf.ProgramSpec `ebpf:"trace___register_chrdev"`
	TraceBpfCheck                            *ebpf.ProgramSpec `ebpf:"trace_bpf_check"`
	TraceCallUsermodehelper                  *ebpf.ProgramSpec `ebpf:"trace_call_usermodehelper"`
	TraceCapCapable                          *ebpf.ProgramSpec `ebpf:"trace_cap_capable"`
	TraceCheckHelperCall                     *ebpf.ProgramSpec `ebpf:"trace_check_helper_call"`
	TraceCheckMapFuncCompatibility           *ebpf.ProgramSpec `ebpf:"trace_check_map_func_compatibility"`
	TraceCommitCreds                         *ebpf.ProgramSpec `ebpf:"trace_commit_creds"`
	TraceDebugfsCreateDir                    *ebpf.ProgramSpec `ebpf:"trace_debugfs_create_dir"`
	TraceDebugfsCreateFile                   *ebpf.ProgramSpec `ebpf:"trace_debugfs_create_file"`
	TraceDeviceAdd                           *ebpf.ProgramSpec `ebpf:"trace_device_add"`
	TraceDoExit                              *ebpf.ProgramSpec `ebpf:"trace_do_exit"`
	TraceDoInitModule                        *ebpf.ProgramSpec `ebpf:"trace_do_init_module"`
	TraceDoMmap                              *ebpf.ProgramSpec `ebpf:"trace_do_mmap"`
	TraceDoSigaction                         *ebpf.ProgramSpec `ebpf:"trace_do_sigaction"`
	TraceDoSplice                            *ebpf.ProgramSpec `ebpf:"trace_do_splice"`
	TraceDoTruncate                          *ebpf.ProgramSpec `ebpf:"trace_do_truncate"`
	TraceExecBinprm                          *ebpf.ProgramSpec `ebpf:"trace_exec_binprm"`
	TraceFdInstall                           *ebpf.ProgramSpec `ebpf:"trace_fd_install"`
	TraceFileModified                        *ebpf.ProgramSpec `ebpf:"trace_file_modified"`
	TraceFileUpdateTime                      *ebpf.ProgramSpec `ebpf:"trace_file_update_time"`
	TraceFilldir64                           *ebpf.ProgramSpec `ebpf:"trace_filldir64"`
	TraceFilpClose                           *ebpf.ProgramSpec `ebpf:"trace_filp_close"`
	TraceInetSockSetState                    *ebpf.ProgramSpec `ebpf:"trace_inet_sock_set_state"`
	TraceInotifyFindInode                    *ebpf.ProgramSpec `ebpf:"trace_inotify_find_inode"`
	TraceKallsymsLookupName                  *ebpf.ProgramSpec `ebpf:"trace_kallsyms_lookup_name"`
	TraceKernelWrite                         *ebpf.ProgramSpec `ebpf:"trace_kernel_write"`
	TraceLoadElfPhdrs                        *ebpf.ProgramSpec `ebpf:"trace_load_elf_phdrs"`
	TraceMmapAlert                           *ebpf.ProgramSpec `ebpf:"trace_mmap_alert"`
	TraceProcCreate                          *ebpf.ProgramSpec `ebpf:"trace_proc_create"`
	TraceRegisterKprobe                      *ebpf.ProgramSpec `ebpf:"trace_register_kprobe"`
	TraceRetRegisterChrdev                   *ebpf.ProgramSpec `ebpf:"trace_ret__register_chrdev"`
	TraceRetDoInitModule                     *ebpf.ProgramSpec `ebpf:"trace_ret_do_init_module"`
	TraceRetDoMmap                           *ebpf.ProgramSpec `ebpf:"trace_ret_do_mmap"`
	TraceRetDoSplice                         *ebpf.ProgramSpec `ebpf:"trace_ret_do_splice"`
	TraceRetExecBinprm                       *ebpf.ProgramSpec `ebpf:"trace_ret_exec_binprm"`
	TraceRetExecBinprm1                      *ebpf.ProgramSpec `ebpf:"trace_ret_exec_binprm1"`
	TraceRetExecBinprm2                      *ebpf.ProgramSpec `ebpf:"trace_ret_exec_binprm2"`
	TraceRetFileModified                     *ebpf.ProgramSpec `ebpf:"trace_ret_file_modified"`
	TraceRetFileUpdateTime                   *ebpf.ProgramSpec `ebpf:"trace_ret_file_update_time"`
	TraceRetInotifyFindInode                 *ebpf.ProgramSpec `ebpf:"trace_ret_inotify_find_inode"`
	TraceRetKallsymsLookupName               *ebpf.ProgramSpec `ebpf:"trace_ret_kallsyms_lookup_name"`
	TraceRetKernelWrite                      *ebpf.ProgramSpec `ebpf:"trace_ret_kernel_write"`
	TraceRetKernelWriteTail                  *ebpf.ProgramSpec `ebpf:"trace_ret_kernel_write_tail"`
	TraceRetLayoutAndAllocate                *ebpf.ProgramSpec `ebpf:"trace_ret_layout_and_allocate"`
	TraceRetRegisterKprobe                   *ebpf.ProgramSpec `ebpf:"trace_ret_register_kprobe"`
	TraceRetSockAllocFile                    *ebpf.ProgramSpec `ebpf:"trace_ret_sock_alloc_file"`
	TraceRetVfsRead                          *ebpf.ProgramSpec `ebpf:"trace_ret_vfs_read"`
	TraceRetVfsReadTail                      *ebpf.ProgramSpec `ebpf:"trace_ret_vfs_read_tail"`
	TraceRetVfsReadv                         *ebpf.ProgramSpec `ebpf:"trace_ret_vfs_readv"`
	TraceRetVfsReadvTail                     *ebpf.ProgramSpec `ebpf:"trace_ret_vfs_readv_tail"`
	TraceRetVfsWrite                         *ebpf.ProgramSpec `ebpf:"trace_ret_vfs_write"`
	TraceRetVfsWriteTail                     *ebpf.ProgramSpec `ebpf:"trace_ret_vfs_write_tail"`
	TraceRetVfsWritev                        *ebpf.ProgramSpec `ebpf:"trace_ret_vfs_writev"`
	TraceRetVfsWritevTail                    *ebpf.ProgramSpec `ebpf:"trace_ret_vfs_writev_tail"`
	TraceSecurityBpfMap                      *ebpf.ProgramSpec `ebpf:"trace_security_bpf_map"`
	TraceSecurityBpfProg                     *ebpf.ProgramSpec `ebpf:"trace_security_bpf_prog"`
	TraceSecurityBprmCheck                   *ebpf.ProgramSpec `ebpf:"trace_security_bprm_check"`
	TraceSecurityFileIoctl                   *ebpf.ProgramSpec `ebpf:"trace_security_file_ioctl"`
	TraceSecurityFileMprotect                *ebpf.ProgramSpec `ebpf:"trace_security_file_mprotect"`
	TraceSecurityFileOpen                    *ebpf.ProgramSpec `ebpf:"trace_security_file_open"`
	TraceSecurityInodeMknod                  *ebpf.ProgramSpec `ebpf:"trace_security_inode_mknod"`
	TraceSecurityInodeRename                 *ebpf.ProgramSpec `ebpf:"trace_security_inode_rename"`
	TraceSecurityInodeSymlink                *ebpf.ProgramSpec `ebpf:"trace_security_inode_symlink"`
	TraceSecurityInodeUnlink                 *ebpf.ProgramSpec `ebpf:"trace_security_inode_unlink"`
	TraceSecurityKernelPostReadFile          *ebpf.ProgramSpec `ebpf:"trace_security_kernel_post_read_file"`
	TraceSecurityKernelReadFile              *ebpf.ProgramSpec `ebpf:"trace_security_kernel_read_file"`
	TraceSecurityMmapFile                    *ebpf.ProgramSpec `ebpf:"trace_security_mmap_file"`
	TraceSecuritySbMount                     *ebpf.ProgramSpec `ebpf:"trace_security_sb_mount"`
	TraceSecuritySkClone                     *ebpf.ProgramSpec `ebpf:"trace_security_sk_clone"`
	TraceSecuritySocketAccept                *ebpf.ProgramSpec `ebpf:"trace_security_socket_accept"`
	TraceSecuritySocketBind                  *ebpf.ProgramSpec `ebpf:"trace_security_socket_bind"`
	TraceSecuritySocketConnect               *ebpf.ProgramSpec `ebpf:"trace_security_socket_connect"`
	TraceSecuritySocketCreate                *ebpf.ProgramSpec `ebpf:"trace_security_socket_create"`
	TraceSecuritySocketListen                *ebpf.ProgramSpec `ebpf:"trace_security_socket_listen"`
	TraceSecuritySocketRecvmsg               *ebpf.ProgramSpec `ebpf:"trace_security_socket_recvmsg"`
	TraceSecuritySocketSendmsg               *ebpf.ProgramSpec `ebpf:"trace_security_socket_sendmsg"`
	TraceSecuritySocketSetsockopt            *ebpf.ProgramSpec `ebpf:"trace_security_socket_setsockopt"`
	TraceSockAllocFile                       *ebpf.ProgramSpec `ebpf:"trace_sock_alloc_file"`
	TraceSwitchTaskNamespaces                *ebpf.ProgramSpec `ebpf:"trace_switch_task_namespaces"`
	TraceSysEnter                            *ebpf.ProgramSpec `ebpf:"trace_sys_enter"`
	TraceSysExit                             *ebpf.ProgramSpec `ebpf:"trace_sys_exit"`
	TraceTracepointProbeRegisterPrioMayExist *ebpf.ProgramSpec `ebpf:"trace_tracepoint_probe_register_prio_may_exist"`
	TraceUtimesCommon                        *ebpf.ProgramSpec `ebpf:"trace_utimes_common"`
	TraceVfsRead                             *ebpf.ProgramSpec `ebpf:"trace_vfs_read"`
	TraceVfsReadv                            *ebpf.ProgramSpec `ebpf:"trace_vfs_readv"`
	TraceVfsUtimes                           *ebpf.ProgramSpec `ebpf:"trace_vfs_utimes"`
	TraceVfsWrite                            *ebpf.ProgramSpec `ebpf:"trace_vfs_write"`
	TraceVfsWritev                           *ebpf.ProgramSpec `ebpf:"trace_vfs_writev"`
	TracepointCgroupCgroupAttachTask         *ebpf.ProgramSpec `ebpf:"tracepoint__cgroup__cgroup_attach_task"`
	TracepointCgroupCgroupMkdir              *ebpf.ProgramSpec `ebpf:"tracepoint__cgroup__cgroup_mkdir"`
	TracepointCgroupCgroupRmdir              *ebpf.ProgramSpec `ebpf:"tracepoint__cgroup__cgroup_rmdir"`
	TracepointModuleModuleFree               *ebpf.ProgramSpec `ebpf:"tracepoint__module__module_free"`
	TracepointModuleModuleLoad               *ebpf.ProgramSpec `ebpf:"tracepoint__module__module_load"`
	TracepointRawSyscallsSysEnter            *ebpf.ProgramSpec `ebpf:"tracepoint__raw_syscalls__sys_enter"`
	TracepointRawSyscallsSysExit             *ebpf.ProgramSpec `ebpf:"tracepoint__raw_syscalls__sys_exit"`
	TracepointSchedSchedProcessExec          *ebpf.ProgramSpec `ebpf:"tracepoint__sched__sched_process_exec"`
	TracepointSchedSchedProcessExit          *ebpf.ProgramSpec `ebpf:"tracepoint__sched__sched_process_exit"`
	TracepointSchedSchedProcessFork          *ebpf.ProgramSpec `ebpf:"tracepoint__sched__sched_process_fork"`
	TracepointSchedSchedProcessFree          *ebpf.ProgramSpec `ebpf:"tracepoint__sched__sched_process_free"`
	TracepointSchedSchedSwitch               *ebpf.ProgramSpec `ebpf:"tracepoint__sched__sched_switch"`
	TracepointTaskTaskRename                 *ebpf.ProgramSpec `ebpf:"tracepoint__task__task_rename"`
	UprobeLkmSeeker                          *ebpf.ProgramSpec `ebpf:"uprobe_lkm_seeker"`
	UprobeLkmSeekerSubmitter                 *ebpf.ProgramSpec `ebpf:"uprobe_lkm_seeker_submitter"`
	UprobeMemDumpTrigger                     *ebpf.ProgramSpec `ebpf:"uprobe_mem_dump_trigger"`
	UprobeSeqOpsTrigger                      *ebpf.ProgramSpec `ebpf:"uprobe_seq_ops_trigger"`
	UprobeSyscallTrigger                     *ebpf.ProgramSpec `ebpf:"uprobe_syscall_trigger"`
}

// tracerMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type tracerMapSpecs struct {
	ArgsMap                 *ebpf.MapSpec `ebpf:"args_map"`
	BinaryFilter            *ebpf.MapSpec `ebpf:"binary_filter"`
	BpfAttachMap            *ebpf.MapSpec `ebpf:"bpf_attach_map"`
	BpfAttachTmpMap         *ebpf.MapSpec `ebpf:"bpf_attach_tmp_map"`
	BpfProgLoadMap          *ebpf.MapSpec `ebpf:"bpf_prog_load_map"`
	Bufs                    *ebpf.MapSpec `ebpf:"bufs"`
	CgroupIdFilter          *ebpf.MapSpec `ebpf:"cgroup_id_filter"`
	CgrpctxmapEg            *ebpf.MapSpec `ebpf:"cgrpctxmap_eg"`
	CgrpctxmapIn            *ebpf.MapSpec `ebpf:"cgrpctxmap_in"`
	CommFilter              *ebpf.MapSpec `ebpf:"comm_filter"`
	ConfigMap               *ebpf.MapSpec `ebpf:"config_map"`
	ContainersMap           *ebpf.MapSpec `ebpf:"containers_map"`
	DebugEvents             *ebpf.MapSpec `ebpf:"debug_events"`
	Entrymap                *ebpf.MapSpec `ebpf:"entrymap"`
	EventDataMap            *ebpf.MapSpec `ebpf:"event_data_map"`
	Events                  *ebpf.MapSpec `ebpf:"events"`
	EventsMap               *ebpf.MapSpec `ebpf:"events_map"`
	FdArgPathMap            *ebpf.MapSpec `ebpf:"fd_arg_path_map"`
	FileModificationMap     *ebpf.MapSpec `ebpf:"file_modification_map"`
	FileReadPathFilter      *ebpf.MapSpec `ebpf:"file_read_path_filter"`
	FileTypeFilter          *ebpf.MapSpec `ebpf:"file_type_filter"`
	FileWritePathFilter     *ebpf.MapSpec `ebpf:"file_write_path_filter"`
	FileWrites              *ebpf.MapSpec `ebpf:"file_writes"`
	IgnoredCgroupsMap       *ebpf.MapSpec `ebpf:"ignored_cgroups_map"`
	Inodemap                *ebpf.MapSpec `ebpf:"inodemap"`
	IoFilePathCacheMap      *ebpf.MapSpec `ebpf:"io_file_path_cache_map"`
	KconfigMap              *ebpf.MapSpec `ebpf:"kconfig_map"`
	KsymbolsMap             *ebpf.MapSpec `ebpf:"ksymbols_map"`
	Logs                    *ebpf.MapSpec `ebpf:"logs"`
	LogsCount               *ebpf.MapSpec `ebpf:"logs_count"`
	MntNsFilter             *ebpf.MapSpec `ebpf:"mnt_ns_filter"`
	ModulesMap              *ebpf.MapSpec `ebpf:"modules_map"`
	NetCapEvents            *ebpf.MapSpec `ebpf:"net_cap_events"`
	NetHeapEvent            *ebpf.MapSpec `ebpf:"net_heap_event"`
	NetHeapScratch          *ebpf.MapSpec `ebpf:"net_heap_scratch"`
	NetconfigMap            *ebpf.MapSpec `ebpf:"netconfig_map"`
	NewModuleMap            *ebpf.MapSpec `ebpf:"new_module_map"`
	OomInfo                 *ebpf.MapSpec `ebpf:"oom_info"`
	PidFilter               *ebpf.MapSpec `ebpf:"pid_filter"`
	PidNsFilter             *ebpf.MapSpec `ebpf:"pid_ns_filter"`
	ProcInfoMap             *ebpf.MapSpec `ebpf:"proc_info_map"`
	ProcessTreeMap          *ebpf.MapSpec `ebpf:"process_tree_map"`
	ProgArray               *ebpf.MapSpec `ebpf:"prog_array"`
	ProgArrayTp             *ebpf.MapSpec `ebpf:"prog_array_tp"`
	RecentDeletedModuleMap  *ebpf.MapSpec `ebpf:"recent_deleted_module_map"`
	RecentInsertedModuleMap *ebpf.MapSpec `ebpf:"recent_inserted_module_map"`
	ScratchMap              *ebpf.MapSpec `ebpf:"scratch_map"`
	SignalDataMap           *ebpf.MapSpec `ebpf:"signal_data_map"`
	Signals                 *ebpf.MapSpec `ebpf:"signals"`
	Sockmap                 *ebpf.MapSpec `ebpf:"sockmap"`
	StackAddresses          *ebpf.MapSpec `ebpf:"stack_addresses"`
	Sys32To64Map            *ebpf.MapSpec `ebpf:"sys_32_to_64_map"`
	SysEnterInitTail        *ebpf.MapSpec `ebpf:"sys_enter_init_tail"`
	SysEnterSubmitTail      *ebpf.MapSpec `ebpf:"sys_enter_submit_tail"`
	SysEnterTails           *ebpf.MapSpec `ebpf:"sys_enter_tails"`
	SysExitInitTail         *ebpf.MapSpec `ebpf:"sys_exit_init_tail"`
	SysExitSubmitTail       *ebpf.MapSpec `ebpf:"sys_exit_submit_tail"`
	SysExitTails            *ebpf.MapSpec `ebpf:"sys_exit_tails"`
	SyscallStatsMap         *ebpf.MapSpec `ebpf:"syscall_stats_map"`
	TaskInfoMap             *ebpf.MapSpec `ebpf:"task_info_map"`
	UidFilter               *ebpf.MapSpec `ebpf:"uid_filter"`
	UtsNsFilter             *ebpf.MapSpec `ebpf:"uts_ns_filter"`
	WalkModTreeQueue        *ebpf.MapSpec `ebpf:"walk_mod_tree_queue"`
}

// tracerObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadTracerObjects or ebpf.CollectionSpec.LoadAndAssign.
type tracerObjects struct {
	tracerPrograms
	tracerMaps
}

func (o *tracerObjects) Close() error {
	return _TracerClose(
		&o.tracerPrograms,
		&o.tracerMaps,
	)
}

// tracerMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadTracerObjects or ebpf.CollectionSpec.LoadAndAssign.
type tracerMaps struct {
	ArgsMap                 *ebpf.Map `ebpf:"args_map"`
	BinaryFilter            *ebpf.Map `ebpf:"binary_filter"`
	BpfAttachMap            *ebpf.Map `ebpf:"bpf_attach_map"`
	BpfAttachTmpMap         *ebpf.Map `ebpf:"bpf_attach_tmp_map"`
	BpfProgLoadMap          *ebpf.Map `ebpf:"bpf_prog_load_map"`
	Bufs                    *ebpf.Map `ebpf:"bufs"`
	CgroupIdFilter          *ebpf.Map `ebpf:"cgroup_id_filter"`
	CgrpctxmapEg            *ebpf.Map `ebpf:"cgrpctxmap_eg"`
	CgrpctxmapIn            *ebpf.Map `ebpf:"cgrpctxmap_in"`
	CommFilter              *ebpf.Map `ebpf:"comm_filter"`
	ConfigMap               *ebpf.Map `ebpf:"config_map"`
	ContainersMap           *ebpf.Map `ebpf:"containers_map"`
	DebugEvents             *ebpf.Map `ebpf:"debug_events"`
	Entrymap                *ebpf.Map `ebpf:"entrymap"`
	EventDataMap            *ebpf.Map `ebpf:"event_data_map"`
	Events                  *ebpf.Map `ebpf:"events"`
	EventsMap               *ebpf.Map `ebpf:"events_map"`
	FdArgPathMap            *ebpf.Map `ebpf:"fd_arg_path_map"`
	FileModificationMap     *ebpf.Map `ebpf:"file_modification_map"`
	FileReadPathFilter      *ebpf.Map `ebpf:"file_read_path_filter"`
	FileTypeFilter          *ebpf.Map `ebpf:"file_type_filter"`
	FileWritePathFilter     *ebpf.Map `ebpf:"file_write_path_filter"`
	FileWrites              *ebpf.Map `ebpf:"file_writes"`
	IgnoredCgroupsMap       *ebpf.Map `ebpf:"ignored_cgroups_map"`
	Inodemap                *ebpf.Map `ebpf:"inodemap"`
	IoFilePathCacheMap      *ebpf.Map `ebpf:"io_file_path_cache_map"`
	KconfigMap              *ebpf.Map `ebpf:"kconfig_map"`
	KsymbolsMap             *ebpf.Map `ebpf:"ksymbols_map"`
	Logs                    *ebpf.Map `ebpf:"logs"`
	LogsCount               *ebpf.Map `ebpf:"logs_count"`
	MntNsFilter             *ebpf.Map `ebpf:"mnt_ns_filter"`
	ModulesMap              *ebpf.Map `ebpf:"modules_map"`
	NetCapEvents            *ebpf.Map `ebpf:"net_cap_events"`
	NetHeapEvent            *ebpf.Map `ebpf:"net_heap_event"`
	NetHeapScratch          *ebpf.Map `ebpf:"net_heap_scratch"`
	NetconfigMap            *ebpf.Map `ebpf:"netconfig_map"`
	NewModuleMap            *ebpf.Map `ebpf:"new_module_map"`
	OomInfo                 *ebpf.Map `ebpf:"oom_info"`
	PidFilter               *ebpf.Map `ebpf:"pid_filter"`
	PidNsFilter             *ebpf.Map `ebpf:"pid_ns_filter"`
	ProcInfoMap             *ebpf.Map `ebpf:"proc_info_map"`
	ProcessTreeMap          *ebpf.Map `ebpf:"process_tree_map"`
	ProgArray               *ebpf.Map `ebpf:"prog_array"`
	ProgArrayTp             *ebpf.Map `ebpf:"prog_array_tp"`
	RecentDeletedModuleMap  *ebpf.Map `ebpf:"recent_deleted_module_map"`
	RecentInsertedModuleMap *ebpf.Map `ebpf:"recent_inserted_module_map"`
	ScratchMap              *ebpf.Map `ebpf:"scratch_map"`
	SignalDataMap           *ebpf.Map `ebpf:"signal_data_map"`
	Signals                 *ebpf.Map `ebpf:"signals"`
	Sockmap                 *ebpf.Map `ebpf:"sockmap"`
	StackAddresses          *ebpf.Map `ebpf:"stack_addresses"`
	Sys32To64Map            *ebpf.Map `ebpf:"sys_32_to_64_map"`
	SysEnterInitTail        *ebpf.Map `ebpf:"sys_enter_init_tail"`
	SysEnterSubmitTail      *ebpf.Map `ebpf:"sys_enter_submit_tail"`
	SysEnterTails           *ebpf.Map `ebpf:"sys_enter_tails"`
	SysExitInitTail         *ebpf.Map `ebpf:"sys_exit_init_tail"`
	SysExitSubmitTail       *ebpf.Map `ebpf:"sys_exit_submit_tail"`
	SysExitTails            *ebpf.Map `ebpf:"sys_exit_tails"`
	SyscallStatsMap         *ebpf.Map `ebpf:"syscall_stats_map"`
	TaskInfoMap             *ebpf.Map `ebpf:"task_info_map"`
	UidFilter               *ebpf.Map `ebpf:"uid_filter"`
	UtsNsFilter             *ebpf.Map `ebpf:"uts_ns_filter"`
	WalkModTreeQueue        *ebpf.Map `ebpf:"walk_mod_tree_queue"`
}

func (m *tracerMaps) Close() error {
	return _TracerClose(
		m.ArgsMap,
		m.BinaryFilter,
		m.BpfAttachMap,
		m.BpfAttachTmpMap,
		m.BpfProgLoadMap,
		m.Bufs,
		m.CgroupIdFilter,
		m.CgrpctxmapEg,
		m.CgrpctxmapIn,
		m.CommFilter,
		m.ConfigMap,
		m.ContainersMap,
		m.DebugEvents,
		m.Entrymap,
		m.EventDataMap,
		m.Events,
		m.EventsMap,
		m.FdArgPathMap,
		m.FileModificationMap,
		m.FileReadPathFilter,
		m.FileTypeFilter,
		m.FileWritePathFilter,
		m.FileWrites,
		m.IgnoredCgroupsMap,
		m.Inodemap,
		m.IoFilePathCacheMap,
		m.KconfigMap,
		m.KsymbolsMap,
		m.Logs,
		m.LogsCount,
		m.MntNsFilter,
		m.ModulesMap,
		m.NetCapEvents,
		m.NetHeapEvent,
		m.NetHeapScratch,
		m.NetconfigMap,
		m.NewModuleMap,
		m.OomInfo,
		m.PidFilter,
		m.PidNsFilter,
		m.ProcInfoMap,
		m.ProcessTreeMap,
		m.ProgArray,
		m.ProgArrayTp,
		m.RecentDeletedModuleMap,
		m.RecentInsertedModuleMap,
		m.ScratchMap,
		m.SignalDataMap,
		m.Signals,
		m.Sockmap,
		m.StackAddresses,
		m.Sys32To64Map,
		m.SysEnterInitTail,
		m.SysEnterSubmitTail,
		m.SysEnterTails,
		m.SysExitInitTail,
		m.SysExitSubmitTail,
		m.SysExitTails,
		m.SyscallStatsMap,
		m.TaskInfoMap,
		m.UidFilter,
		m.UtsNsFilter,
		m.WalkModTreeQueue,
	)
}

// tracerPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadTracerObjects or ebpf.CollectionSpec.LoadAndAssign.
type tracerPrograms struct {
	CgroupBpfRunFilterSkb                    *ebpf.Program `ebpf:"cgroup_bpf_run_filter_skb"`
	CgroupMkdirSignal                        *ebpf.Program `ebpf:"cgroup_mkdir_signal"`
	CgroupRmdirSignal                        *ebpf.Program `ebpf:"cgroup_rmdir_signal"`
	CgroupSkbEgress                          *ebpf.Program `ebpf:"cgroup_skb_egress"`
	CgroupSkbIngress                         *ebpf.Program `ebpf:"cgroup_skb_ingress"`
	LkmSeekerKsetTail                        *ebpf.Program `ebpf:"lkm_seeker_kset_tail"`
	LkmSeekerModTreeTail                     *ebpf.Program `ebpf:"lkm_seeker_mod_tree_tail"`
	LkmSeekerNewModOnlyTail                  *ebpf.Program `ebpf:"lkm_seeker_new_mod_only_tail"`
	LkmSeekerProcTail                        *ebpf.Program `ebpf:"lkm_seeker_proc_tail"`
	OomMarkVictim                            *ebpf.Program `ebpf:"oom_mark_victim"`
	SchedProcessExecEventSubmitTail          *ebpf.Program `ebpf:"sched_process_exec_event_submit_tail"`
	SendBin                                  *ebpf.Program `ebpf:"send_bin"`
	SendBinTp                                *ebpf.Program `ebpf:"send_bin_tp"`
	SysDupExitTail                           *ebpf.Program `ebpf:"sys_dup_exit_tail"`
	SysEnterInit                             *ebpf.Program `ebpf:"sys_enter_init"`
	SysEnterSubmit                           *ebpf.Program `ebpf:"sys_enter_submit"`
	SysExitInit                              *ebpf.Program `ebpf:"sys_exit_init"`
	SysExitSubmit                            *ebpf.Program `ebpf:"sys_exit_submit"`
	SyscallAccept4                           *ebpf.Program `ebpf:"syscall__accept4"`
	SyscallExecve                            *ebpf.Program `ebpf:"syscall__execve"`
	SyscallExecveat                          *ebpf.Program `ebpf:"syscall__execveat"`
	SyscallInitModule                        *ebpf.Program `ebpf:"syscall__init_module"`
	TraceRegisterChrdev                      *ebpf.Program `ebpf:"trace___register_chrdev"`
	TraceBpfCheck                            *ebpf.Program `ebpf:"trace_bpf_check"`
	TraceCallUsermodehelper                  *ebpf.Program `ebpf:"trace_call_usermodehelper"`
	TraceCapCapable                          *ebpf.Program `ebpf:"trace_cap_capable"`
	TraceCheckHelperCall                     *ebpf.Program `ebpf:"trace_check_helper_call"`
	TraceCheckMapFuncCompatibility           *ebpf.Program `ebpf:"trace_check_map_func_compatibility"`
	TraceCommitCreds                         *ebpf.Program `ebpf:"trace_commit_creds"`
	TraceDebugfsCreateDir                    *ebpf.Program `ebpf:"trace_debugfs_create_dir"`
	TraceDebugfsCreateFile                   *ebpf.Program `ebpf:"trace_debugfs_create_file"`
	TraceDeviceAdd                           *ebpf.Program `ebpf:"trace_device_add"`
	TraceDoExit                              *ebpf.Program `ebpf:"trace_do_exit"`
	TraceDoInitModule                        *ebpf.Program `ebpf:"trace_do_init_module"`
	TraceDoMmap                              *ebpf.Program `ebpf:"trace_do_mmap"`
	TraceDoSigaction                         *ebpf.Program `ebpf:"trace_do_sigaction"`
	TraceDoSplice                            *ebpf.Program `ebpf:"trace_do_splice"`
	TraceDoTruncate                          *ebpf.Program `ebpf:"trace_do_truncate"`
	TraceExecBinprm                          *ebpf.Program `ebpf:"trace_exec_binprm"`
	TraceFdInstall                           *ebpf.Program `ebpf:"trace_fd_install"`
	TraceFileModified                        *ebpf.Program `ebpf:"trace_file_modified"`
	TraceFileUpdateTime                      *ebpf.Program `ebpf:"trace_file_update_time"`
	TraceFilldir64                           *ebpf.Program `ebpf:"trace_filldir64"`
	TraceFilpClose                           *ebpf.Program `ebpf:"trace_filp_close"`
	TraceInetSockSetState                    *ebpf.Program `ebpf:"trace_inet_sock_set_state"`
	TraceInotifyFindInode                    *ebpf.Program `ebpf:"trace_inotify_find_inode"`
	TraceKallsymsLookupName                  *ebpf.Program `ebpf:"trace_kallsyms_lookup_name"`
	TraceKernelWrite                         *ebpf.Program `ebpf:"trace_kernel_write"`
	TraceLoadElfPhdrs                        *ebpf.Program `ebpf:"trace_load_elf_phdrs"`
	TraceMmapAlert                           *ebpf.Program `ebpf:"trace_mmap_alert"`
	TraceProcCreate                          *ebpf.Program `ebpf:"trace_proc_create"`
	TraceRegisterKprobe                      *ebpf.Program `ebpf:"trace_register_kprobe"`
	TraceRetRegisterChrdev                   *ebpf.Program `ebpf:"trace_ret__register_chrdev"`
	TraceRetDoInitModule                     *ebpf.Program `ebpf:"trace_ret_do_init_module"`
	TraceRetDoMmap                           *ebpf.Program `ebpf:"trace_ret_do_mmap"`
	TraceRetDoSplice                         *ebpf.Program `ebpf:"trace_ret_do_splice"`
	TraceRetExecBinprm                       *ebpf.Program `ebpf:"trace_ret_exec_binprm"`
	TraceRetExecBinprm1                      *ebpf.Program `ebpf:"trace_ret_exec_binprm1"`
	TraceRetExecBinprm2                      *ebpf.Program `ebpf:"trace_ret_exec_binprm2"`
	TraceRetFileModified                     *ebpf.Program `ebpf:"trace_ret_file_modified"`
	TraceRetFileUpdateTime                   *ebpf.Program `ebpf:"trace_ret_file_update_time"`
	TraceRetInotifyFindInode                 *ebpf.Program `ebpf:"trace_ret_inotify_find_inode"`
	TraceRetKallsymsLookupName               *ebpf.Program `ebpf:"trace_ret_kallsyms_lookup_name"`
	TraceRetKernelWrite                      *ebpf.Program `ebpf:"trace_ret_kernel_write"`
	TraceRetKernelWriteTail                  *ebpf.Program `ebpf:"trace_ret_kernel_write_tail"`
	TraceRetLayoutAndAllocate                *ebpf.Program `ebpf:"trace_ret_layout_and_allocate"`
	TraceRetRegisterKprobe                   *ebpf.Program `ebpf:"trace_ret_register_kprobe"`
	TraceRetSockAllocFile                    *ebpf.Program `ebpf:"trace_ret_sock_alloc_file"`
	TraceRetVfsRead                          *ebpf.Program `ebpf:"trace_ret_vfs_read"`
	TraceRetVfsReadTail                      *ebpf.Program `ebpf:"trace_ret_vfs_read_tail"`
	TraceRetVfsReadv                         *ebpf.Program `ebpf:"trace_ret_vfs_readv"`
	TraceRetVfsReadvTail                     *ebpf.Program `ebpf:"trace_ret_vfs_readv_tail"`
	TraceRetVfsWrite                         *ebpf.Program `ebpf:"trace_ret_vfs_write"`
	TraceRetVfsWriteTail                     *ebpf.Program `ebpf:"trace_ret_vfs_write_tail"`
	TraceRetVfsWritev                        *ebpf.Program `ebpf:"trace_ret_vfs_writev"`
	TraceRetVfsWritevTail                    *ebpf.Program `ebpf:"trace_ret_vfs_writev_tail"`
	TraceSecurityBpfMap                      *ebpf.Program `ebpf:"trace_security_bpf_map"`
	TraceSecurityBpfProg                     *ebpf.Program `ebpf:"trace_security_bpf_prog"`
	TraceSecurityBprmCheck                   *ebpf.Program `ebpf:"trace_security_bprm_check"`
	TraceSecurityFileIoctl                   *ebpf.Program `ebpf:"trace_security_file_ioctl"`
	TraceSecurityFileMprotect                *ebpf.Program `ebpf:"trace_security_file_mprotect"`
	TraceSecurityFileOpen                    *ebpf.Program `ebpf:"trace_security_file_open"`
	TraceSecurityInodeMknod                  *ebpf.Program `ebpf:"trace_security_inode_mknod"`
	TraceSecurityInodeRename                 *ebpf.Program `ebpf:"trace_security_inode_rename"`
	TraceSecurityInodeSymlink                *ebpf.Program `ebpf:"trace_security_inode_symlink"`
	TraceSecurityInodeUnlink                 *ebpf.Program `ebpf:"trace_security_inode_unlink"`
	TraceSecurityKernelPostReadFile          *ebpf.Program `ebpf:"trace_security_kernel_post_read_file"`
	TraceSecurityKernelReadFile              *ebpf.Program `ebpf:"trace_security_kernel_read_file"`
	TraceSecurityMmapFile                    *ebpf.Program `ebpf:"trace_security_mmap_file"`
	TraceSecuritySbMount                     *ebpf.Program `ebpf:"trace_security_sb_mount"`
	TraceSecuritySkClone                     *ebpf.Program `ebpf:"trace_security_sk_clone"`
	TraceSecuritySocketAccept                *ebpf.Program `ebpf:"trace_security_socket_accept"`
	TraceSecuritySocketBind                  *ebpf.Program `ebpf:"trace_security_socket_bind"`
	TraceSecuritySocketConnect               *ebpf.Program `ebpf:"trace_security_socket_connect"`
	TraceSecuritySocketCreate                *ebpf.Program `ebpf:"trace_security_socket_create"`
	TraceSecuritySocketListen                *ebpf.Program `ebpf:"trace_security_socket_listen"`
	TraceSecuritySocketRecvmsg               *ebpf.Program `ebpf:"trace_security_socket_recvmsg"`
	TraceSecuritySocketSendmsg               *ebpf.Program `ebpf:"trace_security_socket_sendmsg"`
	TraceSecuritySocketSetsockopt            *ebpf.Program `ebpf:"trace_security_socket_setsockopt"`
	TraceSockAllocFile                       *ebpf.Program `ebpf:"trace_sock_alloc_file"`
	TraceSwitchTaskNamespaces                *ebpf.Program `ebpf:"trace_switch_task_namespaces"`
	TraceSysEnter                            *ebpf.Program `ebpf:"trace_sys_enter"`
	TraceSysExit                             *ebpf.Program `ebpf:"trace_sys_exit"`
	TraceTracepointProbeRegisterPrioMayExist *ebpf.Program `ebpf:"trace_tracepoint_probe_register_prio_may_exist"`
	TraceUtimesCommon                        *ebpf.Program `ebpf:"trace_utimes_common"`
	TraceVfsRead                             *ebpf.Program `ebpf:"trace_vfs_read"`
	TraceVfsReadv                            *ebpf.Program `ebpf:"trace_vfs_readv"`
	TraceVfsUtimes                           *ebpf.Program `ebpf:"trace_vfs_utimes"`
	TraceVfsWrite                            *ebpf.Program `ebpf:"trace_vfs_write"`
	TraceVfsWritev                           *ebpf.Program `ebpf:"trace_vfs_writev"`
	TracepointCgroupCgroupAttachTask         *ebpf.Program `ebpf:"tracepoint__cgroup__cgroup_attach_task"`
	TracepointCgroupCgroupMkdir              *ebpf.Program `ebpf:"tracepoint__cgroup__cgroup_mkdir"`
	TracepointCgroupCgroupRmdir              *ebpf.Program `ebpf:"tracepoint__cgroup__cgroup_rmdir"`
	TracepointModuleModuleFree               *ebpf.Program `ebpf:"tracepoint__module__module_free"`
	TracepointModuleModuleLoad               *ebpf.Program `ebpf:"tracepoint__module__module_load"`
	TracepointRawSyscallsSysEnter            *ebpf.Program `ebpf:"tracepoint__raw_syscalls__sys_enter"`
	TracepointRawSyscallsSysExit             *ebpf.Program `ebpf:"tracepoint__raw_syscalls__sys_exit"`
	TracepointSchedSchedProcessExec          *ebpf.Program `ebpf:"tracepoint__sched__sched_process_exec"`
	TracepointSchedSchedProcessExit          *ebpf.Program `ebpf:"tracepoint__sched__sched_process_exit"`
	TracepointSchedSchedProcessFork          *ebpf.Program `ebpf:"tracepoint__sched__sched_process_fork"`
	TracepointSchedSchedProcessFree          *ebpf.Program `ebpf:"tracepoint__sched__sched_process_free"`
	TracepointSchedSchedSwitch               *ebpf.Program `ebpf:"tracepoint__sched__sched_switch"`
	TracepointTaskTaskRename                 *ebpf.Program `ebpf:"tracepoint__task__task_rename"`
	UprobeLkmSeeker                          *ebpf.Program `ebpf:"uprobe_lkm_seeker"`
	UprobeLkmSeekerSubmitter                 *ebpf.Program `ebpf:"uprobe_lkm_seeker_submitter"`
	UprobeMemDumpTrigger                     *ebpf.Program `ebpf:"uprobe_mem_dump_trigger"`
	UprobeSeqOpsTrigger                      *ebpf.Program `ebpf:"uprobe_seq_ops_trigger"`
	UprobeSyscallTrigger                     *ebpf.Program `ebpf:"uprobe_syscall_trigger"`
}

func (p *tracerPrograms) Close() error {
	return _TracerClose(
		p.CgroupBpfRunFilterSkb,
		p.CgroupMkdirSignal,
		p.CgroupRmdirSignal,
		p.CgroupSkbEgress,
		p.CgroupSkbIngress,
		p.LkmSeekerKsetTail,
		p.LkmSeekerModTreeTail,
		p.LkmSeekerNewModOnlyTail,
		p.LkmSeekerProcTail,
		p.OomMarkVictim,
		p.SchedProcessExecEventSubmitTail,
		p.SendBin,
		p.SendBinTp,
		p.SysDupExitTail,
		p.SysEnterInit,
		p.SysEnterSubmit,
		p.SysExitInit,
		p.SysExitSubmit,
		p.SyscallAccept4,
		p.SyscallExecve,
		p.SyscallExecveat,
		p.SyscallInitModule,
		p.TraceRegisterChrdev,
		p.TraceBpfCheck,
		p.TraceCallUsermodehelper,
		p.TraceCapCapable,
		p.TraceCheckHelperCall,
		p.TraceCheckMapFuncCompatibility,
		p.TraceCommitCreds,
		p.TraceDebugfsCreateDir,
		p.TraceDebugfsCreateFile,
		p.TraceDeviceAdd,
		p.TraceDoExit,
		p.TraceDoInitModule,
		p.TraceDoMmap,
		p.TraceDoSigaction,
		p.TraceDoSplice,
		p.TraceDoTruncate,
		p.TraceExecBinprm,
		p.TraceFdInstall,
		p.TraceFileModified,
		p.TraceFileUpdateTime,
		p.TraceFilldir64,
		p.TraceFilpClose,
		p.TraceInetSockSetState,
		p.TraceInotifyFindInode,
		p.TraceKallsymsLookupName,
		p.TraceKernelWrite,
		p.TraceLoadElfPhdrs,
		p.TraceMmapAlert,
		p.TraceProcCreate,
		p.TraceRegisterKprobe,
		p.TraceRetRegisterChrdev,
		p.TraceRetDoInitModule,
		p.TraceRetDoMmap,
		p.TraceRetDoSplice,
		p.TraceRetExecBinprm,
		p.TraceRetExecBinprm1,
		p.TraceRetExecBinprm2,
		p.TraceRetFileModified,
		p.TraceRetFileUpdateTime,
		p.TraceRetInotifyFindInode,
		p.TraceRetKallsymsLookupName,
		p.TraceRetKernelWrite,
		p.TraceRetKernelWriteTail,
		p.TraceRetLayoutAndAllocate,
		p.TraceRetRegisterKprobe,
		p.TraceRetSockAllocFile,
		p.TraceRetVfsRead,
		p.TraceRetVfsReadTail,
		p.TraceRetVfsReadv,
		p.TraceRetVfsReadvTail,
		p.TraceRetVfsWrite,
		p.TraceRetVfsWriteTail,
		p.TraceRetVfsWritev,
		p.TraceRetVfsWritevTail,
		p.TraceSecurityBpfMap,
		p.TraceSecurityBpfProg,
		p.TraceSecurityBprmCheck,
		p.TraceSecurityFileIoctl,
		p.TraceSecurityFileMprotect,
		p.TraceSecurityFileOpen,
		p.TraceSecurityInodeMknod,
		p.TraceSecurityInodeRename,
		p.TraceSecurityInodeSymlink,
		p.TraceSecurityInodeUnlink,
		p.TraceSecurityKernelPostReadFile,
		p.TraceSecurityKernelReadFile,
		p.TraceSecurityMmapFile,
		p.TraceSecuritySbMount,
		p.TraceSecuritySkClone,
		p.TraceSecuritySocketAccept,
		p.TraceSecuritySocketBind,
		p.TraceSecuritySocketConnect,
		p.TraceSecuritySocketCreate,
		p.TraceSecuritySocketListen,
		p.TraceSecuritySocketRecvmsg,
		p.TraceSecuritySocketSendmsg,
		p.TraceSecuritySocketSetsockopt,
		p.TraceSockAllocFile,
		p.TraceSwitchTaskNamespaces,
		p.TraceSysEnter,
		p.TraceSysExit,
		p.TraceTracepointProbeRegisterPrioMayExist,
		p.TraceUtimesCommon,
		p.TraceVfsRead,
		p.TraceVfsReadv,
		p.TraceVfsUtimes,
		p.TraceVfsWrite,
		p.TraceVfsWritev,
		p.TracepointCgroupCgroupAttachTask,
		p.TracepointCgroupCgroupMkdir,
		p.TracepointCgroupCgroupRmdir,
		p.TracepointModuleModuleFree,
		p.TracepointModuleModuleLoad,
		p.TracepointRawSyscallsSysEnter,
		p.TracepointRawSyscallsSysExit,
		p.TracepointSchedSchedProcessExec,
		p.TracepointSchedSchedProcessExit,
		p.TracepointSchedSchedProcessFork,
		p.TracepointSchedSchedProcessFree,
		p.TracepointSchedSchedSwitch,
		p.TracepointTaskTaskRename,
		p.UprobeLkmSeeker,
		p.UprobeLkmSeekerSubmitter,
		p.UprobeMemDumpTrigger,
		p.UprobeSeqOpsTrigger,
		p.UprobeSyscallTrigger,
	)
}

func _TracerClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed tracer_bpfel_arm64.o
var _TracerBytes []byte
